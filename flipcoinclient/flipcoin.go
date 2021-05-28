package flipcoinclient

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type CoinFlip string

func (cf CoinFlip) IsHeads() bool {
	return cf == "Heads"
}

func (cf CoinFlip) IsTails() bool {
	return cf == "Tails"
}

type Client struct {
	http.Client
	host string
}

func New(host string) (*Client, error) {
	return &Client{
		Client: http.Client{Timeout: 30 * time.Second},
		host:   host,
	}, nil
}

func (c *Client) TestLuck() (bool, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/flipcoin", c.host), nil)
	if err != nil {
		return false, fmt.Errorf("unable to create a request: %w", err)
	}

	res, err := c.Do(req)
	if err != nil {
		return false, fmt.Errorf("unable to perform a request: %w", err)
	}

	var coinFlip CoinFlip

	err = json.NewDecoder(res.Body).Decode(&coinFlip)
	if err != nil {
		return false, fmt.Errorf("unable to decode response: %w", err)
	}

	return coinFlip.IsHeads(), nil
}
