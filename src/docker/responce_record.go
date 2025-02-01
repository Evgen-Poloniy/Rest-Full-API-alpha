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

func getResponseRecord(w http.ResponseWriter, record Record, tableName string) {
	w.Header().Set("Content-Type", "application/json")

	response := map[string]Record{
		tableName: record,
	}

	responseJSON, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		http.Error(w, "Ошибка формирования JSON", http.StatusInternalServerError)
		return
	}

	w.Write(append(responseJSON, '\n'))
}

type Responses struct {
	Data []Record `json:"data"`
}

func getResponseRecords(w http.ResponseWriter, records []Record, tableName string) {
	w.Header().Set("Content-Type", "application/json")

	response := map[string][]Record{
		tableName: records,
	}

	responseJSON, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		http.Error(w, "Ошибка формирования JSON", http.StatusInternalServerError)
		return
	}

	w.Write(append(responseJSON, '\n'))
}
