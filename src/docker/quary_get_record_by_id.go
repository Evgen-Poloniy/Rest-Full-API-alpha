package main

import (
	"net/http"
)

func ckeckRecordOfUsersByID(w http.ResponseWriter, r *http.Request, table string) (RecordsOfUsers, bool) {
	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		getResponseError(w, http.StatusBadRequest, "Параметр user_id обязателен")
		return RecordsOfUsers{}, false
	}

	var record []RecordsOfUsers = make([]RecordsOfUsers, 1)
	err := db.QueryRow("SELECT * FROM "+table+" WHERE user_id = ?", userID).Scan(&record[0].ID, &record[0].Balance, &record[0].Time)

	if err != nil {
		return RecordsOfUsers{}, false
	}
	return record[0], true
}

func getRecordByID(w http.ResponseWriter, r *http.Request, table string) bool {
	record, exist := ckeckRecordOfUsersByID(w, r, table)

	if !exist {
		return false
	}

	var responce Responses
	responce.Users = []RecordsOfUsers{record}

	getResponseRecord(w, &responce, table)
	return true
}
