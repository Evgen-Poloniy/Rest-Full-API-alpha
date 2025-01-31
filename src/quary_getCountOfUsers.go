package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Count struct {
	Out int `json:"out"`
}

func getCountOfUsers(w http.ResponseWriter, r *http.Request) {
	var count int

	err := db.QueryRow("SELECT COUNT(DISTINCT id) FROM users").Scan(&count)
	if err != nil {
		log.Println("Ошибка запроса к БД:", err)
		http.Error(w, "Ошибка запроса к базе данных", http.StatusInternalServerError)
		return
	}

	response := Count{
		Out: count,
	}

	responseJSON, _ := json.MarshalIndent(response, "", "  ")

	w.Write(append(responseJSON, '\n'))
}
