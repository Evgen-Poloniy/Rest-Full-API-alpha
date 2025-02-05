package main

import (
	"log"
	"net/http"
)

type TransactionParameters struct {
	ID     string
	Type   string
	Amount string
}

func createRecordOfTransaction(w http.ResponseWriter, table string, transactionParameters *TransactionParameters) {
	var complited bool = true
	var record []RecordsOfTransactions = make([]RecordsOfTransactions, 1)
	err := db.QueryRow("INSERT INTO "+table+" (user_id, type_transaction, amount, transaction_time) VALUES (?, ?, ?, CURRENT_TIMESTAMP) RETURNING user_id, transaction_id, type_transaction, amount, transaction_time", transactionParameters.ID, transactionParameters.Type, transactionParameters.Amount).Scan(&record[0].UserID, &record[0].TransactionID, &record[0].Type, &record[0].Amount, &record[0].Time)
	if err != nil {
		getResponseError(w, http.StatusInternalServerError, "Ошибка при создании записи о транзакции")
		complited = false
		return
	}

	var responce Responses
	responce.Transactions = record

	getResponseRecord(w, &responce, table)

	if complited {
		log.Printf("Таблица: %s, Действие: %s, Запрос: %s\n", table, "createRecordOfTransaction", "Успешный")
	} else {
		log.Printf("Таблица: %s, Действие: %s, Запрос: %s\n", table, "createRecordOfTransaction", "Неудачный")
	}
}
