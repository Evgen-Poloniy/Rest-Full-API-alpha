package main

import (
	"log"
	"net/http"

	//"os/user"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	initDB()
	http.HandleFunc("/getRecordByID", getRecordByID)
	http.HandleFunc("/getCountOfUsers", getCountOfUsers)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
