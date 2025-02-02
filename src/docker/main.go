package main

import (
	"log"
	"net/http"

	//"os/user"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	initDB()
	/*
		http.HandleFunc("/checkConnection", checkConnectionHandler)
		http.HandleFunc("/users/createRecord", createRecord)
		http.HandleFunc("/users/deleteRecordById", deleteRecordById)
		http.HandleFunc("/users/getRecordByID", getRecordByID)
		http.HandleFunc("/users/getAllRecords", getAllRecords)
		http.HandleFunc("/transaction/getRecordByID", getRecordByID)
		http.HandleFunc("/transaction/getAllRecords", getAllRecords)
		http.HandleFunc("/users/getCountOfUsers", getCountOfUsers)
	*/
	http.HandleFunc("/", dynamicHandler)
	http.HandleFunc("/users/makeTransaction", makeTransaction)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
