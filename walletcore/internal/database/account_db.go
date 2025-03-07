package database

import (
	"database/sql"

	"github.com.br/devfullcycle/fc-ms-wallet/internal/entity"
)

type AccountDB struct {
	DB *sql.DB
}

func NewAccountDB(db *sql.DB) *AccountDB {
	return &AccountDB{
		DB: db,
	}
}

func (a *AccountDB) FindById(id string) (*entity.Account, error) {
	var account entity.Account
	var client entity.Client
	account.Client = &client
	stmt, err := a.DB.Prepare("select a.id, a.client_id, a.balance, a.created_at, c.id, c.name, c.email, c.created_at from accounts a inner join clients c on c.id = a.client_id where a.id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	row := stmt.QueryRow(id)

	if err := row.Scan(&account.ID, &account.Client.ID, &account.Balance, &account.CreatedAt, &client.ID, &client.Name, &client.Email, &client.CreatedAt); err != nil {
		return nil, err
	}

	return &account, nil
}

func (a *AccountDB) Save(account *entity.Account) error {
	stmt, err := a.DB.Prepare("INSERT INTO accounts (id, client_id, balance, created_at) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(account.ID, account.Client.ID, account.Balance, account.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (a *AccountDB) UpdateBalance(account *entity.Account) error {
	stmt, err := a.DB.Prepare("UPDATE accounts SET balance = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(account.Balance, account.ID)
	if err != nil {
		return err
	}
	return nil
}
