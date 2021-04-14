package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
)

type Transaction struct {
	ID []byte
	Inputs []TxInput
	Outputs[]TxOutput
}

type TxOutput struct {
	Value int
	PubKey string
}

type TxInput struct {
	ID []byte
	Out int
	Sig string
}

func CoinbaseTx(to, data string) *Transaction {
	if data == "" {
		data = fmt.Sprintf("Coins to %s", to)
	}

	txin := TxInput{
		ID:  []byte{},
		Out: -1,
		Sig: data,
	}

	txout := TxOutput{100, to}

	transaction := Transaction{
		ID:      nil,
		Inputs:  []TxInput{txin},
		Outputs: []TxOutput{txout},
	}

	transaction.SetID()

	return &transaction
}

func (tx * Transaction) SetID()  {
	var encoded bytes.Buffer
	var hash[32]byte

	encode := gob.NewEncoder(&encoded)
	err := encode.Encode(tx)
	Handle(err)

	hash = sha256.Sum256(encoded.Bytes())
	tx.ID = hash[:]
}

func (tx * Transaction) IsCoinbase() bool {
	return len(tx.Inputs) == 1 && len(tx.Inputs[0].ID) == 0 && tx.Inputs[0].Out == -1
}

func (input * TxInput) CanUnlock(data string)  bool {
	return input.Sig == data
}

func (output *TxOutput) CanBeUnlocked(data string) bool {
	return output.PubKey == data
}