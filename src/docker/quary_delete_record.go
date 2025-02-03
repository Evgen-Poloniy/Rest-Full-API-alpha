package main

import (
	"net/http"
)

func deleteRecordById(w http.ResponseWriter, r *http.Request, table string) bool {
	user_id := r.URL.Query().Get("user_id")
	if user_id == "" {
		getResponseError(w, 400, "Параметр user_id обязателен")
		return false
	}

	req, _ := http.NewRequest("GET", "/"+table+"/getRecordByID?user_id="+user_id, nil)
	_, exist := ckeckRecordByID(w, req, table)

	if !exist {
		return false
	}

	_, err := db.Exec("DELETE FROM "+table+" users WHERE user_id = ?", user_id)

	if err != nil {
		getResponseError(w, http.StatusInternalServerError, "Ошибка при удалении")
		return false
	}
	return true
}
