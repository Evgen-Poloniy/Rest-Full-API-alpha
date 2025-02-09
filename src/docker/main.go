package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	initDB()
	defer db.Close()
	currencySet = CreateSet()

	http.HandleFunc("/", dynamicHandler)
	http.HandleFunc("/checkConnection", checkConnectionHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
