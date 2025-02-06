package main

import (
	"log"
	"net/http"
)

func closeDB(w http.ResponseWriter, r *http.Request) {
	var password string = r.URL.Query().Get("password")
	var ip string = r.RemoteAddr

	if password == "" {
		w.WriteHeader(http.StatusBadRequest)
	}

	if password == dbCfg.Password {
		db.Close()
		w.WriteHeader(http.StatusOK)
		log.Printf("Действие: closeDB, IP: %s, Запрос: Успешный", ip)
		log.Println("Закрыта база данных MySQL")
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		log.Printf("Действие: closeDB, IP: %s, Запрос: Неудачный", ip)
	}
}
