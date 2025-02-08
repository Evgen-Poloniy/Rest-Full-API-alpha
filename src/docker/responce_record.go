package main

import (
	"encoding/json"
	"math"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type RecordsOfUsers struct {
	ID       int     `json:"user_id"`
	Balance  float64 `json:"balance"`
	Currency string  `json:"currency"`
	Time     string  `json:"transaction_time"`
}

type RecordsOfTransactions struct {
	TransactionID int     `json:"transaction_id"`
	UserID        int     `json:"user_id"`
	Type          string  `json:"type_transaction"`
	Amount        float64 `json:"amount"`
	Currency      string  `json:"currency"`
	Time          string  `json:"transaction_time"`
}

type RecordsOfErrors struct {
	Error   int    `json:"error"`
	Message string `json:"message"`
}

type RecordOfAgregation struct {
	Agregation int `json:"count"`
}

type Responses struct {
	Users        []RecordsOfUsers        `json:"users"`
	Transactions []RecordsOfTransactions `json:"transactions"`
	Errors       []RecordsOfErrors       `json:"errors"`
	Agregation   []RecordOfAgregation    `json:"agregation"`
}

func getResponseRecord(w http.ResponseWriter, responce *Responses, tableName string, currency string) bool {
	w.Header().Set("Content-Type", "application/json")

	var records interface{}
	if tableName == "users" {
		for i := range responce.Users {
			responce.Users[i].Currency = currency
			responce.Users[i].Balance = math.Round(responce.Users[i].Balance/currencySet[currency]*100) / 100
		}
		records = responce.Users
	} else if tableName == "transactions" {
		for i := range responce.Transactions {
			responce.Transactions[i].Currency = currency
			responce.Transactions[i].Amount = math.Round(responce.Transactions[i].Amount/currencySet[currency]*100) / 100
		}
		records = responce.Transactions
	} else {
		getResponseError(w, http.StatusInternalServerError, "Ошибка при создании записи")
		return false
	}

	var response interface{}
	switch rec := records.(type) {
	case []RecordsOfUsers:
		if len(rec) == 1 {
			response = map[string]RecordsOfUsers{
				tableName: rec[0],
			}
		} else {
			response = map[string][]RecordsOfUsers{
				tableName: rec,
			}
		}
	case []RecordsOfTransactions:
		if len(rec) == 1 {
			response = map[string]RecordsOfTransactions{
				tableName: rec[0],
			}
		} else {
			response = map[string][]RecordsOfTransactions{
				tableName: rec,
			}
		}
	}

	responseJSON, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		getResponseError(w, http.StatusInternalServerError, "Ошибка формирования JSON")
		return false
	}

	w.Write(append(responseJSON, '\n'))
	return true
}

func getResponseError(w http.ResponseWriter, statusCode int, errMsg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := map[string]RecordsOfErrors{
		"errors": {
			Error:   statusCode,
			Message: errMsg,
		},
	}

	responseJSON, _ := json.MarshalIndent(response, "", "  ")

	w.Write(append(responseJSON, '\n'))
}
