package blockchain

import (
	"bytes"
	"github.com/josefelipetto/golang-blockchain/wallet"
)

type TxOutput struct {
	Value  int
	PubKeyHash []byte
}

type TxInput struct {
	ID  []byte
	Out int
	Signature []byte
	PubKey []byte
}

func (input * TxInput) UsesKey(pubKeyHash []byte) bool {
	lockingHash := wallet.PublicKeyHash(input.PubKey)

	return bytes.Compare(lockingHash, pubKeyHash) == 0
}

func (output *TxOutput) Lock(address []byte) {
	pubKeyHash := wallet.Base58Decode(address)
	pubKeyHash = pubKeyHash[1: len(pubKeyHash) - 4]
	output.PubKeyHash = pubKeyHash
}

func (output *TxOutput) IsLockedWithKey(pubKeyHash []byte) bool {
	return bytes.Compare(output.PubKeyHash, pubKeyHash) == 0
}

func NewTxOutput(value int, address string) *TxOutput {
	txo := &TxOutput{
		Value:      value,
		PubKeyHash: nil,
	}

	txo.Lock([]byte(address))

	return txo
}
