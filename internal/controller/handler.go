package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/banking-system/middleware"
)

type TransferRequest struct {
	Sender    string `json:"sender"`
	Recipient string `json:"recipient"`
	Amount    int32  `json:"amount"`
}

type TransferResponse struct {
	Sender       string `json:"sender"`
	IsSuccess    bool   `json:"isSuccess"`
	ErrorMessage string `json:"errorMessage"`
}

func (c *Controller) Transfer(w http.ResponseWriter, r *http.Request) {
	var treq *TransferRequest
	err := json.NewDecoder(r.Body).Decode(&treq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Println(c.transactionService.Transfer())

	tresp := &TransferResponse{
		IsSuccess:    true,
		Sender:       treq.Sender,
		ErrorMessage: "no error",
	}
	middleware.JSON(w, http.StatusOK, tresp)
}
