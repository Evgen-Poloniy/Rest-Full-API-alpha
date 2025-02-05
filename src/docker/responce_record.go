package main

import (
	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

// Структура для хранения записи
type RecordsOfUsers struct {
	ID      int     `json:"user_id"`
	Balance float64 `json:"balance"`
	Time    string  `json:"transaction_time"`
}

type RecordsOfTransactions struct {
	TransactionID int     `json:"transaction_id"`
	UserID        int     `json:"user_id"`
	Type          string  `json:"type_transaction"`
	Amount        float64 `json:"balance"`
	Time          string  `json:"transaction_time"`
}

type Responses struct {
	Users        []RecordsOfUsers        `json:"users"`
	Transactions []RecordsOfTransactions `json:"transactions"`
}

func getResponseRecord(w http.ResponseWriter, responce *Responses, tableName string) {
	w.Header().Set("Content-Type", "application/json")

	var records interface{}
	if tableName == "users" {
		records = responce.Users
	} else if tableName == "transactions" {
		records = responce.Transactions
	} else {
		getResponseError(w, http.StatusInternalServerError, "Ошибка при создании записи")
		return
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
		return
	}

	w.Write(append(responseJSON, '\n'))
}
