package dbclient

import (
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"

	"github.com/lambda-direct/gocast-wager/env"
)

type FlipResult struct {
	Id       uint64 `json:"id"`
	Username string `json:"username"`
	Amount   uint16 `json:"amount"`
	Result   bool   `json:"result"`
}

type Client struct {
	db *pg.DB
}

func New(dbConfig env.DB) (*Client, error) {
	db := pg.Connect(&pg.Options{
		User:     dbConfig.Username,
		Password: dbConfig.Password,
		Addr:     fmt.Sprintf("%s:%d", dbConfig.Host, dbConfig.Port),
		Database: dbConfig.Name,
	})

	models := []interface{}{
		(*FlipResult)(nil),
	}

	for _, model := range models {
		if err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			IfNotExists: true,
		}); err != nil {
			return nil, err
		}
	}

	return &Client{db}, nil
}

func (c *Client) SaveResult(result *FlipResult) error {
	_, err := c.db.Model(result).Insert()
	return err
}

func (c *Client) Close() error {
	return c.db.Close()
}
