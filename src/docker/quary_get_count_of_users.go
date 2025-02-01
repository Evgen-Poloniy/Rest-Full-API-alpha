package main

import (
	"encoding/json"
	"net/http"
)

type Count struct {
	Out int `json:"out"`
}

func getResponceCount(w http.ResponseWriter, count Count) {
	w.Header().Set("Content-Type", "application/json")

	responseJSON, err := json.MarshalIndent(count, "", "  ")
	if err != nil {
		http.Error(w, "Ошибка формирования JSON", http.StatusInternalServerError)
		return
	}

	w.Write(append(responseJSON, '\n'))
}

func getCountOfUsers(w http.ResponseWriter, r *http.Request) {
	var count int

	err := db.QueryRow("SELECT COUNT(DISTINCT id) FROM users").Scan(&count)
	if err != nil {
		getResponseError(w, 204, "В таблице нет пользователей")
		return
	}

	response := Count{
		Out: count,
	}

	getResponceCount(w, response)
}
