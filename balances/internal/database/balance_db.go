package database

import (
	"database/sql"
	"errors"

	"github.com.br/devfullcycle/fc-ms-balances/internal/entity"
)

type BalanceDB struct {
	DB *sql.DB
}

func NewBalanceDB(db *sql.DB) *BalanceDB {
	return &BalanceDB{
		DB: db,
	}
}

func (c *BalanceDB) Get(accountId string) (*entity.Balance, error) {
	balance := &entity.Balance{}
	stmt, err := c.DB.Prepare("select id, name, account_id, balance FROM balances WHERE account_id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	row := stmt.QueryRow(accountId)

	if err := row.Scan(&balance.ID, &balance.Name, &balance.AccountID, &balance.Balance); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("no balance found for this account: " + accountId)
		} else {
			return nil, err
		}
	}

	return balance, nil
}

func (c *BalanceDB) Save(balance *entity.Balance) error {
	stmt, err := c.DB.Prepare("INSERT INTO balances (name, account_id, balance) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(balance.Name, balance.AccountID, balance.Balance)
	if err != nil {
		return err
	}
	return nil
}

func (c *BalanceDB) Update(balance *entity.Balance) error {
	stmt, err := c.DB.Prepare("UPDATE balances SET balance = ? WHERE account_id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(balance.Balance, balance.AccountID)
	if err != nil {
		return err
	}
	return nil
}
