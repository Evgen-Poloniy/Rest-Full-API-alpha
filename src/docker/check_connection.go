package main

import (
	"fmt"
	"net/http"
)

func checkConnectionHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	ip := r.RemoteAddr
	fmt.Printf("Действие: checkConnection c ip: %s", ip)
}
