package parser

import (
	"strconv"
	"strings"
	"time"

	"github.com/lucavallin/transit/pkg/transaction"
)

// Ing keeps ing bank data
type Ing struct{}

// NewIng creates a new Ing struct
func NewIng() Ing {
	return Ing{}
}

// Parse parses Ing struct data into a collection of transactions
func (o Ing) Parse(records [][]string) (transaction.Collection, error) {

	var transactions transaction.Collection

	for _, record := range records {
		// Init transaction and append to slice
		transaction := &transaction.Transaction{
			Date:      o.parseDate(record[0]),
			Name:      record[1],
			Account:   record[2],
			ToAccount: record[3],
			Code:      record[4],
			Direction: o.parseDirection(record[5]),
			Amount:    o.parseAmount(record[6]),
			Type:      record[7],
			Notes:     record[8],
		}
		transactions = append(transactions, transaction)
	}

	return transactions, nil

}

// Converts csv string to date
func (o Ing) parseDate(rawDate string) time.Time {
	// Is there a way to unhardcode this golang value?
	date, err := time.Parse("20060102", rawDate)
	if err != nil {
		return time.Time{}
	}

	return date
}

// Converts csv string to float64
func (o Ing) parseAmount(rawAmount string) float64 {
	amountSanitized := strings.Replace(rawAmount, ",", ".", -1)
	amount, _ := strconv.ParseFloat(amountSanitized, 64)

	return amount
}

// Converts csv direction to transaction.Direction
func (o Ing) parseDirection(rawDirection string) int8 {
	direction := transaction.Incoming

	// Most transactions are outgoing, saves a check
	if rawDirection == "Af" {
		direction = transaction.Outgoing
	}

	return direction
}
