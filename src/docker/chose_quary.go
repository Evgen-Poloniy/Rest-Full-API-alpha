package main

import (
	"fmt"
	"net/http"
	"strings"
)

func dynamicHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")

	if len(parts) < 3 {
		getResponseError(w, http.StatusBadRequest, "Неверный запрос")
		return
	}

	table := parts[1]
	action := parts[2]

	var completed bool = false

	switch action {
	case "createRecord":
		completed = createRecord(w, r, table)
	case "deleteRecordByID":
		completed = deleteRecordById(w, r, table)
	case "getRecordByID":
		completed = getRecordByID(w, r, table)
	case "getAllRecords":
		completed = getAllRecords(w, r, table)
	case "getCountOfRecords":
		completed = getCountOfRecords(w, table)
	default:
		getResponseError(w, http.StatusNotFound, "Неизвестное действие")
	}

	var status string
	if completed {
		status = "Успешный"
	} else {
		status = "Неудачный"
	}
	fmt.Printf("Таблица: %s, Действие: %s, Запрос: %s\n", table, action, status)
}
