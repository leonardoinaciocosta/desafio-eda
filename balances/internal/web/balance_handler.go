package web

import (
	"encoding/json"
	"fmt"
	"net/http"

	checkbalance "github.com.br/devfullcycle/fc-ms-balances/internal/usecase/check_balance"
)

type WebCheckBalanceHandler struct {
	CheckBalanceUseCase checkbalance.CheckBalanceUseCase
}

func NewWebCheckBalanceHandler(checkBalanceUseCase checkbalance.CheckBalanceUseCase) *WebCheckBalanceHandler {
	return &WebCheckBalanceHandler{
		CheckBalanceUseCase: checkBalanceUseCase,
	}
}

func (h *WebCheckBalanceHandler) CheckBalance(w http.ResponseWriter, r *http.Request) {
	var account_id = r.PathValue("account_id")

	if account_id == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("account_id not supplied"))
		return
	}

	fmt.Println(account_id)
	output, err := h.CheckBalanceUseCase.Execute(r.PathValue("account_id"))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		fmt.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(output)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		fmt.Println(err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
