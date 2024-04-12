package checkbalance

import (
	"github.com.br/devfullcycle/fc-ms-balances/internal/gateway"
)

type CheckBalanceInputDTO struct {
	AccountID string
}

type CheckBalanceOutputDTO struct {
	Name      string
	AccountID string
	Balance   float64
}

type CheckBalanceUseCase struct {
	BalanceGateway gateway.BalanceGateway
}

func NewCheckBalanceUseCase(balanceGateway gateway.BalanceGateway) *CheckBalanceUseCase {
	return &CheckBalanceUseCase{
		BalanceGateway: balanceGateway,
	}
}

func (u *CheckBalanceUseCase) Execute(accountID string) (*CheckBalanceOutputDTO, error) {
	balance, err := u.BalanceGateway.Get(accountID)

	if err != nil {
		return nil, err
	}

	return &CheckBalanceOutputDTO{
		Name:      balance.Name,
		AccountID: balance.AccountID,
		Balance:   balance.Balance,
	}, nil
}
