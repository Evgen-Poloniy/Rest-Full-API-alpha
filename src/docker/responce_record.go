package main

import (
	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

// Структура для хранения записи
type Record struct {
	ID      int     `json:"id"`
	Balance float64 `json:"balance"`
	Time    string  `json:"last_time"`
}

func getResponseRecord(w http.ResponseWriter, record Record) {
	w.Header().Set("Content-Type", "application/json")

	responseJSON, err := json.MarshalIndent(record, "", "  ")
	if err != nil {
		http.Error(w, "Ошибка формирования JSON", http.StatusInternalServerError)
		return
	}

	w.Write(append(responseJSON, '\n'))
}

func getResponseRecords(w http.ResponseWriter, records []Record) {
	w.Header().Set("Content-Type", "application/json")

	// Формируем JSON с отступами
	responseJSON, err := json.MarshalIndent(records, "", "  ")
	if err != nil {
		http.Error(w, "Ошибка формирования JSON", http.StatusInternalServerError)
		return
	}

	// Отправляем JSON клиенту
	w.Write(append(responseJSON, '\n'))
}
