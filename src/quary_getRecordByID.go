package main

import (
	"net/http"
)

// Получение записи по ID
func getRecordByID(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("id")
	if userID == "" {
		responseError(w, 400, "Параметр id обязателен")
		return
	}

	var record Record
	err := db.QueryRow("SELECT id, balance, last_time FROM users WHERE id = ?", userID).Scan(&record.ID, &record.Balance, &record.Time)

	if err != nil {
		responseError(w, 404, "Пользователь не найден")
		return
	}

	responseRecord(w, record)
}
