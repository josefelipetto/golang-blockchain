package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

type BlockChain struct {
	blocks []*Block
}

type Block struct {
	Hash []byte
	Data []byte
	PrevHash []byte
}

func (b *Block) DeriveHash()  {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash}
	block.DeriveHash()
	return block
}

func (c *BlockChain) AddBlock(data string)  {
	prevBlock := c.blocks[len(c.blocks) - 1]
	newBlock := CreateBlock(data, prevBlock.Hash)
	c.blocks = append(c.blocks, newBlock)
}

func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

func InitBlockChain()  *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}

func main() {
	chain := InitBlockChain()

	chain.AddBlock("First block after Genesis")
	chain.AddBlock("Second block after Genesis")
	chain.AddBlock("Third block after Genesis")

	for _, block := range chain.blocks {
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println("=====================================")
	}
}
