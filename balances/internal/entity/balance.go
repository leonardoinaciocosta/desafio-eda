package entity

import (
	"errors"

	"github.com/google/uuid"
)

type Balance struct {
	ID        string
	Name      string
	AccountID string
	Balance   float64
}

func NewBalance(name string, accountID string, balance float64) (*Balance, error) {
	client := &Balance{
		ID:        uuid.NewString(),
		Name:      name,
		AccountID: accountID,
		Balance:   balance,
	}

	err := client.Validate()

	if err != nil {
		return nil, err
	}

	return client, nil
}

func (c *Balance) Validate() error {
	if c.Name == "" {
		return errors.New("Name is required")
	}

	if c.AccountID == "" {
		return errors.New("Account ID is required")
	}

	return nil
}

func (c *Balance) Update(name string, accountID string, balance float64) error {
	c.Name = name
	c.AccountID = accountID
	c.Balance = balance
	err := c.Validate()

	if err != nil {
		return err
	}

	return nil
}
