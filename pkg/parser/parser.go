package parser

import "github.com/lucavallin/transit/pkg/transaction"

// Parser is an interface for parsers
type Parser interface {
	Parse(records [][]string) (transaction.Collection, error)
}
