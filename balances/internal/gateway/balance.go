package gateway

import "github.com.br/devfullcycle/fc-ms-balances/internal/entity"

type BalanceGateway interface {
	Get(accountId string) (*entity.Balance, error)
	Save(client *entity.Balance) error
	Update(client *entity.Balance) error
}
