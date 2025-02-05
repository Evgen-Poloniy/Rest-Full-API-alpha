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

/*
func createRecordOfTransaction(w http.ResponseWriter, table string, transactionParameters *TransactionParameters) {
	var record []RecordsOfTransactions = make([]RecordsOfTransactions, 1)
	err := db.QueryRow("INSERT INTO "+table+" (user_id, type_transaction, amount, transaction_time) VALUES (?, ?, ?, CURRENT_TIMESTAMP) RETURNING user_id, transaction_id, type_transaction, amount, transaction_time", transactionParameters.ID, transactionParameters.Type, transactionParameters.Amount)
	if err != nil {
		getResponseError(w, http.StatusInternalServerError, "Ошибка при создании записи о транзакции")
		log.Printf("Таблица: %s, Действие: %s, Запрос: %s\n", table, "createRecordOfTransaction", "Неудачный")
		return
	}

	var responce Responses
	responce.Transactions = record

	getResponseRecord(w, &responce, table)

	log.Printf("Таблица: %s, Действие: %s, Запрос: %s\n", table, "createRecordOfTransaction", "Успешный")
}
*/

func createRecordOfTransaction(w http.ResponseWriter, table string, transactionParameters *TransactionParameters) {
	result, err := db.Exec("INSERT INTO "+table+" (user_id, type_transaction, amount, transaction_time) VALUES (?, ?, ?, CURRENT_TIMESTAMP)",
		transactionParameters.ID, transactionParameters.Type, transactionParameters.Amount)

	if err != nil {
		log.Printf("Ошибка при создании записи в таблице %s: %v", table, err)
		getResponseError(w, http.StatusInternalServerError, "Ошибка при создании записи о транзакции")
		return
	}

	transactionID, err := result.LastInsertId()
	if err != nil {
		log.Printf("Ошибка получения ID последней транзакции: %v", err)
		getResponseError(w, http.StatusInternalServerError, "Ошибка при получении ID транзакции")
		return
	}

	var record RecordsOfTransactions
	err = db.QueryRow("SELECT user_id, transaction_id, type_transaction, amount, transaction_time FROM "+table+" WHERE transaction_id = ?", transactionID).
		Scan(&record.UserID, &record.TransactionID, &record.Type, &record.Amount, &record.Time)

	if err != nil {
		log.Printf("Ошибка получения данных транзакции: %v", err)
		getResponseError(w, http.StatusInternalServerError, "Ошибка при получении данных транзакции")
		return
	}

	var response Responses
	response.Transactions = []RecordsOfTransactions{record}
	getResponseRecord(w, &response, table)

	log.Printf("Таблица: %s, Действие: createRecordOfTransaction, Запрос: Успешный", table)
}
