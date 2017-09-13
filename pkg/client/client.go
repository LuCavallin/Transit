package client

// Client is an interface for providers
type Client interface {
	Fetch() ([][]string, error)
}
