package main

import (
	"log"
	"net/http"
)

func checkConnectionHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	ip := r.RemoteAddr
	log.Println("Действие: checkConnection c ip: %s", ip)
}
