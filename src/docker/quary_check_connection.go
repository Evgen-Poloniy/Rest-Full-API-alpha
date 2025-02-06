package main

import (
	"log"
	"net/http"
)

func checkConnectionHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	var ip string = r.RemoteAddr
	log.Printf("Действие: checkConnection, IP: %s", ip)
}
