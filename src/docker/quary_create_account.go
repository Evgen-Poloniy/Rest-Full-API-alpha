package main

import (
	"net/http"
	"strconv"
)

func createRecord(w http.ResponseWriter, r *http.Request) {
	balanceStr := r.URL.Query().Get("balance")
	if balanceStr == "" {
		getResponseError(w, http.StatusBadRequest, "Не указан баланс")
		return
	}

	balance, err := strconv.ParseFloat(balanceStr, 64)
	if err != nil {
		getResponseError(w, http.StatusBadRequest, "Некорректный формат баланса")
		return
	}

	_, err = db.Exec("INSERT INTO users (balance) VALUES (?)", balance)
	if err != nil {
		getResponseError(w, http.StatusInternalServerError, "Ошибка при вставке записи")
		return
	}

	var record Record
	err = db.QueryRow("SELECT * FROM users WHERE id = (SELECT MAX(id) FROM users);").Scan(&record.ID, &record.Balance, &record.Time)

	if err != nil {
		getResponseError(w, 404, "Пользователь не найден")
		return
	}

	getResponseRecord(w, record, "users")
}
