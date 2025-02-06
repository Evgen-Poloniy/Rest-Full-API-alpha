package main

import (
	"net/http"
	"strconv"
)

func createRecordOfUsers(w http.ResponseWriter, r *http.Request, table string) (bool, bool) {
	balanceStr := r.URL.Query().Get("balance")
	if balanceStr == "" {
		getResponseError(w, http.StatusBadRequest, "Не указан баланс")
		return false, false
	}

	balance, err := strconv.ParseFloat(balanceStr, 64)
	if err != nil {
		getResponseError(w, http.StatusBadRequest, "Некорректный формат баланса")
		return false, false
	}

	_, err = db.Exec("INSERT INTO "+table+" (balance, transaction_time) VALUES (?, CURRENT_TIMESTAMP)", balance)
	if err != nil {
		getResponseError(w, http.StatusInternalServerError, "Ошибка при создании записи о пользователе")
		return false, false
	}

	var record []RecordsOfUsers = make([]RecordsOfUsers, 1)
	err = db.QueryRow("SELECT * FROM "+table+" WHERE user_id = (SELECT MAX(user_id) FROM users);").Scan(&record[0].ID, &record[0].Balance, &record[0].Time)

	if err != nil {
		getResponseError(w, http.StatusNotFound, "Пользователь не найден")
		return false, false
	}

	var responce Responses
	responce.Users = record

	return true, getResponseRecord(w, &responce, table)
}
