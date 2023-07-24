package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type Block struct {
	nonce        int
	previousHash [32]byte
	timestamp    float64
	transactions []*Transaction
}

func NewBlock(nonce int, previousHash [32]byte, transactions []*Transaction) *Block {
	b := new(Block)
	b.nonce = nonce
	b.previousHash = previousHash
	b.timestamp = float64(time.Now().UnixNano())
	b.transactions = transactions
	return b
}

func (b *Block) Print() {
	fmt.Printf("%20s\t%d\n", "nonce", b.nonce)
	fmt.Printf("%20s\t%x\n", "previousHash", b.previousHash)
	fmt.Printf("%20s\t%f\n", "timestamp", b.timestamp)
	for _, value := range b.transactions {
		fmt.Printf("%s%s Transaction%s\n", strings.Repeat(" ", 23), strings.Repeat("-", 23), strings.Repeat("-", 23))
		value.Print()
	}
}

func (b *Block) toJSON() ([]byte, error) {

	// convert object block into a JSON compatible block
	JSONBlock := struct {
		Nonce        int            `json:"nonce"`
		PreviousHash [32]byte       `json:"previousHash"`
		Timestamp    float64        `json:"timestamp"`
		Transactions []*Transaction `json:"transactions"`
	}{
		Nonce:        b.nonce,
		PreviousHash: b.previousHash,
		Timestamp:    b.timestamp,
		Transactions: b.transactions,
	}

	return json.Marshal(JSONBlock)
}

func (b *Block) CalculateBlockHash() [32]byte {
	// convert the current block to JSON
	tmpJSON, _ := b.toJSON()
	// calculate the sum256 of the JSON (it's unique)
	return sha256.Sum256([]byte(tmpJSON))
}
