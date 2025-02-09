package main

import (
	"encoding/json"
	"net/http"
)

func getResponceCount(w http.ResponseWriter, count RecordOfAgregation) {
	w.Header().Set("Content-Type", "application/json")

	response := map[string]RecordOfAgregation{
		"agregation": count,
	}

	responseJSON, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		http.Error(w, "Ошибка формирования JSON", http.StatusInternalServerError)
		return
	}

	w.Write(append(responseJSON, '\n'))
}

func getCountOfRecords(w http.ResponseWriter, table string) bool {
	var count int

	err := db.QueryRow("SELECT COUNT(DISTINCT user_id) FROM ?", table).Scan(&count)
	if err != nil {
		getResponseError(w, 204, "В таблице нет пользователей")
		return false
	}

	response := RecordOfAgregation{
		Agregation: count,
	}

	getResponceCount(w, response)
	return true
}
