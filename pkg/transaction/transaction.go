package transaction

import (
	"time"
)

// Transaction data
type Transaction struct {
	Date      time.Time
	Name      string
	Account   string
	ToAccount string
	Code      string
	Direction int8
	Amount    float64
	Type      string
	Notes     string
}

const (
	// Incoming defines outgoing transactions
	Incoming int8 = 1
	// Outgoing defines incoming transactions
	Outgoing int8 = -1
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

// ReportByName returns the total amount grouped by name of the receiver/sender
func (c Collection) ReportByName(direction int8) map[string]float64 {
	groupedTransactions := make(map[string]float64)

	transactions := c.getByDirection(direction)
	for _, transaction := range transactions {
		groupedTransactions[transaction.Name] += transaction.Amount
	}

	return groupedTransactions
}

// getByDirection filters transactions based on their direction
func (c Collection) getByDirection(direction int8) Collection {
	var filteredTransactions Collection
	for _, current := range c {
		if current.Direction == direction {
			filteredTransactions = append(filteredTransactions, current)
		}
	}

	return filteredTransactions
}

// getByName filters transactions by name of the external guy (not you, basically)
func (c Collection) getByName(name string) Collection {
	var filteredTransactions Collection
	for _, current := range c {
		if current.Name == name {
			filteredTransactions = append(filteredTransactions, current)
		}
	}

	return filteredTransactions
}
