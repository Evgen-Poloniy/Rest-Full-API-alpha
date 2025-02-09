package main

import (
	"database/sql"
	"net/http"
	"strconv"
)

func getAllRecords(w http.ResponseWriter, r *http.Request, table string) bool {
	var currency string = r.URL.Query().Get("currency")
	var parametr string = r.URL.Query().Get("parametr")
	var order string = r.URL.Query().Get("order")
	var limit string = r.URL.Query().Get("limit")
	var offsetStr string = r.URL.Query().Get("offset")

	if order == "" {
		order = "ASC"
	}

	if offsetStr == "" {
		offsetStr = "0"
	}

	offset, errParse := strconv.ParseInt(offsetStr, 10, 64)
	if errParse != nil {
		getResponseError(w, http.StatusBadRequest, "Некорректное значение offset")
		return false
	}

	var balance float64 = 0
	if !updateExchangeRates(w, &balance, &currency) {
		return false
	}

	if table == "users" {
		if parametr == "" {
			parametr = "user_id"
		} else if parametr == "id" {
			parametr = "user_id"
		} else if parametr == "amount" {
			parametr = "balance"
		}
	} else {
		if parametr == "" {
			parametr = "transaction_id"
		} else if parametr == "id" {
			parametr = "transaction_id"
		}
	}

	var rows *sql.Rows
	var err error

	if limit == "" {
		rows, err = db.Query("SELECT * FROM " + table + " ORDER BY " + parametr + " " + order)
	} else {
		rows, err = db.Query("SELECT * FROM "+table+" ORDER BY "+parametr+" "+order+" LIMIT ? OFFSET ?", limit, offset)
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

	return getResponseRecord(w, &responce, table, currency)
}
