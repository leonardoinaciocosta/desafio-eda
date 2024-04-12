package web

import (
	"encoding/json"
	"fmt"
	"net/http"

	createclient "github.com.br/devfullcycle/fc-ms-wallet/internal/usecase/create_client"
	listclient "github.com.br/devfullcycle/fc-ms-wallet/internal/usecase/list_client"
)

type WebClientHandler struct {
	CreateClientUseCase createclient.CreateClientUseCase
	ListClientUseCase   listclient.ListClientUseCase
}

func NewWebClientHandler(createClientUseCase createclient.CreateClientUseCase, listClientUseCase listclient.ListClientUseCase) *WebClientHandler {
	return &WebClientHandler{
		CreateClientUseCase: createClientUseCase,
		ListClientUseCase:   listClientUseCase,
	}
}

func (h *WebClientHandler) ProcessRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		h.ListClient(w, r)
	} else {
		h.CreateClient(w, r)
	}
}

func (h *WebClientHandler) CreateClient(w http.ResponseWriter, r *http.Request) {
	var dto createclient.CreteClientInputDTO
	err := json.NewDecoder(r.Body).Decode(&dto)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
		return
	}

	output, err := h.CreateClientUseCase.Execute(dto)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		fmt.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(output)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		fmt.Println(err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *WebClientHandler) ListClient(w http.ResponseWriter, r *http.Request) {
	output, err := h.ListClientUseCase.Execute()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		fmt.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(output)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		fmt.Println(err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
