package main

import (
	"net/http"
)

func ckeckRecordByID(w http.ResponseWriter, r *http.Request, table string) (Record, bool) {
	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		getResponseError(w, 400, "Параметр user_id обязателен")
		return Record{}, false
	}

	var record []Record = make([]Record, 1)
	err := db.QueryRow("SELECT * FROM "+table+" WHERE user_id = ?", userID).Scan(&record[0].ID, &record[0].Balance, &record[0].Time)

	if err != nil {
		getResponseError(w, 404, "Пользователь не найден")
		return Record{}, false
	}
	return record[0], true
}

func getRecordByID(w http.ResponseWriter, r *http.Request, table string) bool {
	record, exist := ckeckRecordByID(w, r, table)

	if !exist {
		return false
	}

	getResponseRecord(w, []Record{record}, table)
	return true
}
