package main

import (
	"fmt"
	"log"
	"net/http"
)

func openDB(w http.ResponseWriter, r *http.Request) {
	var password string = r.URL.Query().Get("password")
	var ip string = r.RemoteAddr

	if password == "" {
		w.WriteHeader(http.StatusBadRequest)
	}

	if password == dbCfg.Password {
		var err error = db.Ping()
		if err != nil {
			fmt.Fprintf(w, "База данных уже работает")
			return
		}
		initDB()
		w.WriteHeader(http.StatusOK)
		log.Printf("Действие: openDB, IP: %s, Запрос: Успешный", ip)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		log.Printf("Действие: openDB, IP: %s, Запрос: Неудачный", ip)
	}
}
