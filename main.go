package main

import (
	"flag"
	"fmt"
	"github.com/josefelipetto/golang-blockchain/blockchain"
	"log"
	"os"
	"runtime"
	"strconv"
)

type CommandLine struct {}

func (cli * CommandLine) printUsage() {
	fmt.Println("Usage:")
	fmt.Println(" print - Prints the blocks in the chain")
	fmt.Println(" getbalance -address ADDRESS - get balance for address")
	fmt.Println(" createblockchain -address ADDRESS creates a blockchain")
	fmt.Println(" send -from FROM -to TO -amount AMOUNT - Send an amount from FROM to TO")

}

func (cli * CommandLine) validateArgs()  {
	if len(os.Args) < 2 {
		cli.printUsage()
		runtime.Goexit()
	}
}

func (cli * CommandLine) printChain()  {
	chain := blockchain.ContinueBlockChain("")
	defer chain.Database.Close()
	iter := chain.Iterator()

	for {
		block := iter.Next()
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println("=====================================")

		pow := blockchain.NewProof(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()

		if len(block.PrevHash) == 0 {
			break
		}
	}
}

func (cli * CommandLine) run()  {
	cli.validateArgs()

	printChainCmd := flag.NewFlagSet("print", flag.ExitOnError)
	getBalanceCmd := flag.NewFlagSet("getbalance", flag.ExitOnError)
	createBlockchainCmd := flag.NewFlagSet("createblockchain", flag.ExitOnError)
	sendCmd := flag.NewFlagSet("send", flag.ExitOnError)

	getBalanceAddress := getBalanceCmd.String("address", "", "The address")
	createBlockChainAddress := createBlockchainCmd.String("address", "", "The address")
	sendFrom := sendCmd.String("from", "", "Source wallet address")
	sendTo := sendCmd.String("to", "", "Destination wallet address")
	sendAmount := sendCmd.Int("amount", 0, "Amount to send")

	switch os.Args[1] {
	case "print":
		err := printChainCmd.Parse(os.Args[2:])
		blockchain.Handle(err)
	case "getbalance":
		err := getBalanceCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "createblockchain":
		err := createBlockchainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "send":
		err := sendCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}

	default:
		cli.printUsage()
		runtime.Goexit()
	}

	if printChainCmd.Parsed() {
		cli.printChain()
	}

	if getBalanceCmd.Parsed() {
		if *getBalanceAddress == "" {
			getBalanceCmd.Usage()
			runtime.Goexit()
		}
		cli.getBalance(*getBalanceAddress)
	}

	if createBlockchainCmd.Parsed() {
		if *createBlockChainAddress == "" {
			createBlockchainCmd.Usage()
			runtime.Goexit()
		}
		cli.createBlockChain(*createBlockChainAddress)
	}

	if sendCmd.Parsed() {
		if *sendFrom == "" || *sendTo == "" || *sendAmount < 0 {
			sendCmd.Usage()
			runtime.Goexit()
		}
		cli.send(*sendFrom, *sendTo, *sendAmount)
	}
 }

func (cli * CommandLine) createBlockChain(address string)  {
	chain := blockchain.InitBlockChain(address)
	chain.Database.Close()
	fmt.Println("Finished!")
}

func (cli * CommandLine) getBalance(address string) {
	chain := blockchain.ContinueBlockChain(address)
	defer chain.Database.Close()

	balance := 0
	UTXOs := chain.FindUTXO(address)

	for _, out := range UTXOs {
		balance += out.Value
	}

	fmt.Printf("Balance of %s: %d\n", address, balance)
}

func (cli * CommandLine) send(from, to string, amount int) {
	chain := blockchain.ContinueBlockChain(from)
	defer chain.Database.Close()

	tx := blockchain.NewTransaction(from, to, amount, chain)
	chain.AddBlock([]*blockchain.Transaction{tx})
	fmt.Println("Sent successfully")
}


func main() {
	defer os.Exit(0)
	cli := CommandLine{}
	cli.run()
}
