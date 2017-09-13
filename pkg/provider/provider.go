package provider

import (
	"github.com/lucavallin/transit/pkg/client"
	"github.com/lucavallin/transit/pkg/parser"
	"github.com/lucavallin/transit/pkg/transaction"
)

// Provider groups client and parser to retrieve transactions
type Provider struct {
	client client.Client
	parser parser.Parser
}

// NewProvider returns new provider
func NewProvider(client client.Client, parser parser.Parser) Provider {
	return Provider{client, parser}
}

// Transactions reads and parse data to provide a collection of transactions
func (o Provider) Transactions() (transaction.Collection, error) {
	content, err := o.client.Fetch()
	if err != nil {
		return nil, err
	}

	return o.parser.Parse(content)
}
