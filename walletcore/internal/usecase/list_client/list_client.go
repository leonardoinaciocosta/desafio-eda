package listclient

import (
	"github.com.br/devfullcycle/fc-ms-wallet/internal/entity"
	"github.com.br/devfullcycle/fc-ms-wallet/internal/gateway"
)

type ListClientOutputDTO struct {
	Clients []entity.Client
}

type ListClientUseCase struct {
	ClientGateway gateway.ClientGateway
}

func NewListClientUseCase(clientGateway gateway.ClientGateway) *ListClientUseCase {
	return &ListClientUseCase{
		ClientGateway: clientGateway,
	}
}

func (u *ListClientUseCase) Execute() (*ListClientOutputDTO, error) {
	var lista, err = u.ClientGateway.List()

	if err != nil {
		return nil, err
	}

	return &ListClientOutputDTO{
		Clients: lista,
	}, nil
}
