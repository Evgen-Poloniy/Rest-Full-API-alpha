package main

import (
	"net/http"
)

func ckeckRecordOfUsersByID(w http.ResponseWriter, r *http.Request, table string) (interface{}, bool) {
	if table == "users" {
		var userID string = r.URL.Query().Get("user_id")
		if userID == "" {
			getResponseError(w, http.StatusBadRequest, "Параметр user_id обязателен")
			return RecordsOfUsers{}, false
		}

		var record []RecordsOfUsers = make([]RecordsOfUsers, 1)
		err := db.QueryRow("SELECT * FROM "+table+" WHERE user_id = ?", userID).Scan(&record[0].ID, &record[0].Balance, &record[0].Time)

		if err != nil {
			return RecordsOfUsers{}, false
		}
		return record[0], true
	} else {
		var userID string = r.URL.Query().Get("transaction_id")
		if userID == "" {
			getResponseError(w, http.StatusBadRequest, "Параметр transaction_id обязателен")
			return nil, false
		}

		var record []RecordsOfTransactions = make([]RecordsOfTransactions, 1)
		err := db.QueryRow("SELECT * FROM "+table+" WHERE transaction_id = ?", userID).Scan(&record[0].TransactionID, &record[0].UserID, &record[0].Type, &record[0].Amount, &record[0].Time)

		if err != nil {
			return RecordsOfTransactions{}, false
		}
		return record[0], true
	}
}

func getRecordByID(w http.ResponseWriter, r *http.Request, table string) bool {
	record, exist := ckeckRecordOfUsersByID(w, r, table)

	var currency string = r.URL.Query().Get("currency")

	var balance float64 = 0
	if !updateExchangeRates(w, &balance, &currency) {
		return false
	}

	if !exist {
		if table == "users" {
			getResponseError(w, http.StatusNotFound, "Пользователь не найден")
		} else {
			getResponseError(w, http.StatusNotFound, "Транзакция не найдена")
		}
		return false
	}

	var responce Responses
	if table == "users" {
		responce.Users = []RecordsOfUsers{record.(RecordsOfUsers)}
	} else {
		responce.Transactions = []RecordsOfTransactions{record.(RecordsOfTransactions)}
	}

	return getResponseRecord(w, &responce, table, currency)
}
