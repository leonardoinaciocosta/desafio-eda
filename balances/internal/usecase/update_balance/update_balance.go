package updatebalance

import (
	"github.com.br/devfullcycle/fc-ms-balances/internal/entity"
	"github.com.br/devfullcycle/fc-ms-balances/internal/gateway"
)

type UpdateBalanceInputDTO struct {
	Name      string
	AccountID string
	Balance   float64
}

type UpdateBalanceUseCase struct {
	BalanceGateway gateway.BalanceGateway
}

func NewUpdateBalanceUseCase(balanceGateway gateway.BalanceGateway) *UpdateBalanceUseCase {
	return &UpdateBalanceUseCase{
		BalanceGateway: balanceGateway,
	}
}

func (u *UpdateBalanceUseCase) Execute(input UpdateBalanceInputDTO) error {
	balance, err := entity.NewBalance(input.Name, input.AccountID, input.Balance)

	if err != nil {
		return err
	}

	err = u.BalanceGateway.Save(balance)

	if err != nil {
		return err
	}

	return nil
}
