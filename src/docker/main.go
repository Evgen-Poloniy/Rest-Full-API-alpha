package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	initDB()
	defer db.Close()

	http.HandleFunc("/", dynamicHandler)
	http.HandleFunc("/checkConnection", checkConnectionHandler)
	http.HandleFunc("/closeDB", closeDB)
	http.HandleFunc("/openDB", openDB)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
