package main

import (
	"log"
	"net/http"
	"strconv"
)

type TransactionParameters struct {
	ID     string
	Type   string
	Amount string
}

func createRecordOfTransaction(w http.ResponseWriter, table string, transactionParameters *TransactionParameters, currency *string) bool {
	amount, err := strconv.ParseFloat(transactionParameters.Amount, 64)
	if err != nil {
		getResponseError(w, http.StatusBadRequest, "Неверный формат суммы транзакции")
		return false
	}

	amount *= currencySet[*currency]

	result, err := db.Exec("INSERT INTO "+table+" (user_id, type_transaction, amount, transaction_time) VALUES (?, ?, ?, CURRENT_TIMESTAMP)",
		transactionParameters.ID, transactionParameters.Type, amount)

	if err != nil {
		log.Printf("Ошибка при создании записи в таблице %s: %v", table, err)
		getResponseError(w, http.StatusInternalServerError, "Ошибка при создании записи о транзакции")
		return false
	}

	transactionID, err := result.LastInsertId()
	if err != nil {
		log.Printf("Ошибка получения ID последней транзакции: %v", err)
		getResponseError(w, http.StatusInternalServerError, "Ошибка при получении ID транзакции")
		return false
	}

	var record RecordsOfTransactions
	err = db.QueryRow("SELECT user_id, transaction_id, type_transaction, amount, transaction_time FROM "+table+" WHERE transaction_id = ?", transactionID).
		Scan(&record.UserID, &record.TransactionID, &record.Type, &record.Amount, &record.Time)

	if err != nil {
		log.Printf("Ошибка получения данных транзакции: %v", err)
		getResponseError(w, http.StatusInternalServerError, "Ошибка при получении данных транзакции")
		return false
	}

	var response Responses
	response.Transactions = []RecordsOfTransactions{record}
	getResponseRecord(w, &response, table, *currency)

	return true
}
