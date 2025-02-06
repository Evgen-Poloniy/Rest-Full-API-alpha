package main

import (
	"database/sql"
	"net/http"
)

func getAllRecords(w http.ResponseWriter, r *http.Request, table string) bool {
	limitStr := r.URL.Query().Get("limit")
	var rows *sql.Rows
	var err error

	if table == "users" {
		if limitStr != "" {
			rows, err = db.Query("SELECT * FROM "+table+" ORDER BY user_id LIMIT ?", limitStr)
		} else {
			rows, err = db.Query("SELECT * FROM " + table + " ORDER BY user_id")
		}
	} else {
		if limitStr != "" {
			rows, err = db.Query("SELECT * FROM "+table+" ORDER BY transaction_id LIMIT ?", limitStr)
		} else {
			rows, err = db.Query("SELECT * FROM " + table + " ORDER BY transaction_id")
		}
	}

	if err != nil {
		getResponseError(w, http.StatusInternalServerError, "Ошибка запроса к БД")
		return false
	}
	defer rows.Close()

	var responce Responses
	if table == "users" {
		var records []RecordsOfUsers

		for rows.Next() {
			var record RecordsOfUsers
			if err := rows.Scan(&record.ID, &record.Balance, &record.Time); err != nil {
				getResponseError(w, http.StatusInternalServerError, "Ошибка чтения данных")
				return false
			}
			records = append(records, record)
		}

		if len(records) == 0 {
			getResponseError(w, http.StatusNotFound, "В таблице нет данных")
			return false
		}

		responce.Users = records
	} else {
		var records []RecordsOfTransactions
		for rows.Next() {
			var record RecordsOfTransactions
			if err := rows.Scan(&record.TransactionID, &record.UserID, &record.Type, &record.Amount, &record.Time); err != nil {
				getResponseError(w, http.StatusInternalServerError, "Ошибка чтения данных")
				return false
			}
			records = append(records, record)
		}

		if len(records) == 0 {
			getResponseError(w, http.StatusNotFound, "В таблице нет данных")
			return false
		}

		responce.Transactions = records
	}
	getResponseRecord(w, &responce, table)

	return true
}
