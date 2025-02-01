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
}
