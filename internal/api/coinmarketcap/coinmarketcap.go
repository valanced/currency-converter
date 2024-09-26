package coinmarketcap

import (
	"context"
)

type Client struct {
	apiKey string
}

func NewClient(apiKey string) *Client {
	return &Client{apiKey: apiKey}
}

func (c *Client) FetchRate(ctx context.Context, from, to string) (float64, error) {
	panic("implement me")
}
