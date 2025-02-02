package main

import (
	"net/http"
)

// Получение записи по ID
func getRecordByID(w http.ResponseWriter, r *http.Request, table string) {
	userID := r.URL.Query().Get("id")
	if userID == "" {
		getResponseError(w, 400, "Параметр id обязателен")
		return
	}

	var record Record
	err := db.QueryRow("SELECT * FROM "+table+" WHERE id = ?", userID).Scan(&record.ID, &record.Balance, &record.Time)

	if err != nil {
		getResponseError(w, 404, "Пользователь не найден")
		return
	}

	getResponseRecord(w, record, "users")
}
