package main

import (
	"encoding/json"
	"fmt"
)

const (
	maxTransactionsPerBlock = 2
)

type Transaction struct {
	sender    string
	recipient string
	value     float32
}

func NewTransaction(s string, r string, v float32) *Transaction {
	return &Transaction{s, r, v}
}

func (t *Transaction) Print() {
	fmt.Printf("%35s\t%s\n", "Sender: ", t.sender)
	fmt.Printf("%35s\t%s\n", "Recipient: ", t.recipient)
	fmt.Printf("%35s\t%.10f\n", "Value: ", t.value)
}

func (b *Transaction) toJSON() ([]byte, error) {

	// convert object transaction into a JSON compatible transaction
	JSONBlock := struct {
		Sender    string  `json:"sender"`
		Recipient string  `json:"recipient"`
		Value     float32 `json:"value"`
	}{
		Sender:    b.sender,
		Recipient: b.recipient,
		Value:     b.value,
	}

	return json.Marshal(JSONBlock)
}
