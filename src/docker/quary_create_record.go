package main

import (
	"net/http"
	"strconv"
)

func createRecord(w http.ResponseWriter, r *http.Request, table string) bool {
	balanceStr := r.URL.Query().Get("balance")
	if balanceStr == "" {
		getResponseError(w, http.StatusBadRequest, "Не указан баланс")
		return false
	}

	balance, err := strconv.ParseFloat(balanceStr, 64)
	if err != nil {
		getResponseError(w, http.StatusBadRequest, "Некорректный формат баланса")
		return false
	}

	_, err = db.Exec("INSERT INTO "+table+" (balance) VALUES (?)", balance)
	if err != nil {
		getResponseError(w, http.StatusInternalServerError, "Ошибка при вставке записи")
		return false
	}

	var record []Record = make([]Record, 1)
	err = db.QueryRow("SELECT * FROM "+table+" WHERE user_id = (SELECT MAX(user_id) FROM users);").Scan(&record[0].ID, &record[0].Balance, &record[0].Time)

	if err != nil {
		getResponseError(w, 404, "Пользователь не найден")
		return false
	}

	getResponseRecord(w, record, table)
	return true
}
