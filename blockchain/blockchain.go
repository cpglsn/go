package main

import (
	"crypto/sha256"
	"fmt"
	"strings"
)

const (
	difficult = 3
)

type BlockChain struct {
	transactionPool []*Transaction
	chain           []*Block
}

func NewBlockChain() *BlockChain {
	bc := new(BlockChain)
	// create first block with hash = 0
	tmpHash := sha256.Sum256([]byte("0"))
	// append the created block as the first one of the chain converting []byte in string
	bc.chain = append(bc.chain, NewBlock(0, tmpHash, nil))
	return bc
}

func (bc *BlockChain) AddBlock(nonce int, transactions []*Transaction) {
	// take the last created block
	previousBlock := bc.chain[len(bc.chain)-1]
	// append the block to the chain
	bc.chain = append(bc.chain, NewBlock(nonce, previousBlock.CalculateBlockHash(), transactions))
}

func (bc *BlockChain) Print() {
	for i, block := range bc.chain {
		fmt.Printf("%s Block %3d %s\n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		block.Print()
	}
	fmt.Println(strings.Repeat("=", 61))

	// print buffer transactions
	fmt.Printf("%s Buffered Transactions %s\n", strings.Repeat("-", 30), strings.Repeat("-", 30))

	for _, value := range bc.transactionPool {
		fmt.Printf("Sender: %s\n", value.sender)
		fmt.Printf("Recipient: %s\n", value.recipient)
		fmt.Printf("Value: %0.1f\n", value.value)
	}

	fmt.Println(strings.Repeat("=", 61))
}

func (bc *BlockChain) addTransaction(t *Transaction) {
	bc.transactionPool = append(bc.transactionPool, t)
	if len(bc.transactionPool) >= maxTransactionsPerBlock {
		bc.emptyTransactionPool()
	}
}

func (bc *BlockChain) emptyTransactionPool() {
	// create a temporary transaction slice to avoid alias
	tmpTransactionPool := make([]*Transaction, len(bc.transactionPool))
	copy(tmpTransactionPool, bc.transactionPool)

	// create the new block
	bc.AddBlock(9, tmpTransactionPool)

	// clean transaction buffer
	bc.transactionPool = nil
}

func (bc *BlockChain) ValidateBlock(b Block) bool {
	if b.nonce[:difficult] == strings.Repeat("0", difficult) &&
	   b. {

	}
}
