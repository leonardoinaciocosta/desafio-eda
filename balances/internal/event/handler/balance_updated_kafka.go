package handler

import (
	"encoding/json"
	"fmt"

	"github.com.br/devfullcycle/fc-ms-balances/internal/event"
	updatebalance "github.com.br/devfullcycle/fc-ms-balances/internal/usecase/update_balance"
)

type UpdateBalanceKafkaHandler struct {
	balanceUpdated       *event.BalanceUpdated
	updateBalanceUseCase *updatebalance.UpdateBalanceUseCase
}

type BalanceUpdatedDTO struct {
	AccountIDFrom        string  `json:"account_id_from"`
	AccountIDTo          string  `json:"account_id_to"`
	BalanceAccountIDFrom float64 `json:"balance_account_id_from"`
	BalanceAccountIDTo   float64 `json:"balance_account_id_to"`
}

func NewUpdateBalanceKafkaHandler(balanceUpdated *event.BalanceUpdated, updateBalanceUseCase *updatebalance.UpdateBalanceUseCase) *UpdateBalanceKafkaHandler {
	return &UpdateBalanceKafkaHandler{
		balanceUpdated:       balanceUpdated,
		updateBalanceUseCase: updateBalanceUseCase,
	}
}

func (h *UpdateBalanceKafkaHandler) Handle() error {
	fmt.Println("UpdateBalanceKafkaHandler called")
	jsonString, err2 := json.Marshal(h.balanceUpdated.Payload)
	if err2 != nil {
		fmt.Println("UpdateBalanceKafkaHandler Error: ", err2)
		return err2
	}
	balanceUpdatedDTO := BalanceUpdatedDTO{}
	err := json.Unmarshal(jsonString, &balanceUpdatedDTO)
	if err != nil {
		fmt.Println("UpdateBalanceKafkaHandler Error: ", err)
		return err
	}

	balanceFrom, err := h.updateBalanceUseCase.BalanceGateway.Get(balanceUpdatedDTO.AccountIDFrom)
	if err != nil {
		fmt.Println("UpdateBalanceKafkaHandler Error: ", err)
		return err
	}

	balanceFrom.Balance = balanceUpdatedDTO.BalanceAccountIDFrom
	err = h.updateBalanceUseCase.BalanceGateway.Update(balanceFrom)
	if err != nil {
		fmt.Println("UpdateBalanceKafkaHandler Error: ", err)
		return err
	}

	balanceTo, err := h.updateBalanceUseCase.BalanceGateway.Get(balanceUpdatedDTO.AccountIDTo)
	if err != nil {
		fmt.Println("UpdateBalanceKafkaHandler Error: ", err)
		return err
	}

	balanceTo.Balance = balanceUpdatedDTO.BalanceAccountIDTo
	err = h.updateBalanceUseCase.BalanceGateway.Update(balanceTo)
	if err != nil {
		fmt.Println("UpdateBalanceKafkaHandler Error: ", err)
		return err
	}

	fmt.Println("UpdateBalanceKafkaHandler completed")
	return nil
}
