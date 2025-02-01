package main

import (
	"log"
	"net/http"

	//"os/user"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	initDB()
	http.HandleFunc("/createRecord", createRecord)
	http.HandleFunc("/getCountOfUsers", getCountOfUsers)
	http.HandleFunc("/getRecordByID", getRecordByID)
	http.HandleFunc("/getAllRecords", getAllRecords)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
