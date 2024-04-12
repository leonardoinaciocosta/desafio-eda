package database

import (
	"database/sql"

	"github.com.br/devfullcycle/fc-ms-wallet/internal/entity"
)

type ClientDB struct {
	DB *sql.DB
}

func NewClientDB(db *sql.DB) *ClientDB {
	return &ClientDB{
		DB: db,
	}
}

func (c *ClientDB) Get(id string) (*entity.Client, error) {
	client := &entity.Client{}
	stmt, err := c.DB.Prepare("select id, name, email, created_at FROM clients WHERE id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	row := stmt.QueryRow(id)

	if err := row.Scan(&client.ID, &client.Name, &client.Email, &client.CreatedAt); err != nil {
		return nil, err
	}

	return client, nil
}

func (c *ClientDB) Save(client *entity.Client) error {
	stmt, err := c.DB.Prepare("INSERT INTO clients (id, name, email, created_at) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(client.ID, client.Name, client.Email, client.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (c *ClientDB) List() ([]entity.Client, error) {
	var clients []entity.Client
	rows, err := c.DB.Query("select id, name, email, created_at from clients")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var client entity.Client

		if err := rows.Scan(&client.ID, &client.Name, &client.Email, &client.CreatedAt); err != nil {
			return nil, err
		}

		stmt, err := c.DB.Prepare("select a.id, a.balance, a.created_at from accounts a where a.client_id = ?")
		if err != nil {
			return nil, err
		}

		defer stmt.Close()
		rowsAccount, err := stmt.Query(client.ID)

		if err != nil {
			return nil, err
		}

		for rowsAccount.Next() {
			var account entity.Account

			if err := rowsAccount.Scan(&account.ID, &account.Balance, &account.CreatedAt); err != nil {
				return nil, err
			}

			client.Accounts = append(client.Accounts, &account)
		}

		clients = append(clients, client)
	}

	if err = rows.Err(); err != nil {
		return clients, err
	}

	return clients, nil
}
