package main

import (
	"fmt"
	"os"

	"github.com/lucavallin/transit/pkg/parser/csv"

	"github.com/davecgh/go-spew/spew"
	"github.com/lucavallin/transit/pkg/client"
	"github.com/lucavallin/transit/pkg/provider"
	"github.com/lucavallin/transit/pkg/transaction"
)

func main() {
	var path string

	// Checks on input
	if len(os.Args) > 1 {
		path = os.Args[1]
	} else {
		fmt.Println("No filepath provided.")
		os.Exit(1)
	}

	// Intro
	fmt.Print("\n===| Transit Home Bookkeeping |===\n\n")

	// Fetch and parse data
	provider := provider.NewProvider(client.NewCsv(path), parser.NewIng())
	transactions, err := provider.Transactions()
	if err != nil {
		fmt.Printf("Could not process data.")
		os.Exit(1)
	}

	spew.Dump(transactions)

	fmt.Printf("Incoming transactions total: %f\n", transactions.GetTotalAmount(transaction.DirectionIn))
	fmt.Printf("Outgoing transactions total: %f\n", transactions.GetTotalAmount(transaction.DirectionOut))
}
