package main

import (
	"fmt"
	"os"

	"github.com/lucavallin/transit/pkg/parser/csv"

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

	fmt.Printf("> Analyzing outgoing transactions...\n\n")
	for name, amount := range transactions.ReportByName(transaction.Outgoing) {
		fmt.Printf("%s: %.2f\n", name, amount)
	}
	fmt.Printf("\nTotal: %.2f\n", transactions.GetTotalAmount(transaction.Outgoing))

	fmt.Println()
	fmt.Println()

	fmt.Printf("> Analyzing incoming transactions...\n\n")
	for name, amount := range transactions.ReportByName(transaction.Incoming) {
		fmt.Printf("%s: %.2f\n", name, amount)
	}
	fmt.Printf("\nTotal: %.2f\n", transactions.GetTotalAmount(transaction.Incoming))
}
