package gateway

import "github.com.br/devfullcycle/fc-ms-wallet/internal/entity"

type ClientGateway interface {
	List() ([]entity.Client, error)
	Get(id string) (*entity.Client, error)
	Save(client *entity.Client) error
}
