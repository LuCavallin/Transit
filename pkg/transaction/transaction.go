package transaction

import (
	"time"
)

// Transaction data
type Transaction struct {
	Date        time.Time
	Description string
	Account     string
	ToAccount   string
	Code        string
	Direction   int8
	Amount      float64
	Type        string
	Notes       string
}

const (
	// DirectionOut defines outgoing transactions
	DirectionOut int8 = 1
	// DirectionIn defines incoming transactions
	DirectionIn int8 = -1
)

// Collection contains Transactions
type Collection []*Transaction

// GetTotalAmount returns the total amount in transactions Slice
func (c Collection) GetTotalAmount(direction int8) float64 {
	var totalAmount float64

	for _, transaction := range c.getByDirection(direction) {
		totalAmount += transaction.Amount
	}

	return totalAmount
}

// GroupByName returns the total amount grouped by name of the receiver/sender
func (c Collection) GroupByName(direction int8) float64 {
	var totalAmount float64

	for _, transaction := range c.getByDirection(direction) {
		totalAmount += transaction.Amount
	}

	return totalAmount
}

// getByDirection filters transactions from a slice based on their direction
func (c Collection) getByDirection(direction int8) Collection {
	var filteredTransactions Collection
	for _, transaction := range c {
		if transaction.Direction == direction {
			filteredTransactions = append(filteredTransactions, transaction)
		}
	}

	return filteredTransactions
}
