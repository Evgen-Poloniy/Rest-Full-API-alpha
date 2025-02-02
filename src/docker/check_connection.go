package main

import (
	"net/http"
)

func checkConnectionHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
