package main

import (
	"net/http"
)

func deleteRecordById(w http.ResponseWriter, r *http.Request, table string) {
	id := r.URL.Query().Get("id")
	if id == "" {
		getResponseError(w, 400, "Параметр id обязателен")
		return
	}

	_, err := db.Exec("DELETE FROM "+table+" users WHERE id = ?", id)
	if err != nil {
		getResponseError(w, http.StatusNotFound, "Запись с указанным id не найдена")
		return
	}
}
