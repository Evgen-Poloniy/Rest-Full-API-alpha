package main

import (
	"log"
	"net/http"
	"strings"
)

func dynamicHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")

	if len(parts) < 3 {

		getResponseError(w, http.StatusBadRequest, "Неверный запрос")
		return
	}

	var table string = parts[1]
	var action string = parts[2]
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
	case "makeTransaction":
		completed = makeTransaction(w, r)
	default:
		getResponseError(w, http.StatusNotFound, "Неизвестное действие")
	}

	var status string
	if completed {
		status = "Успешный"
	} else {
		status = "Неудачный"
	}
	log.Println("Таблица: %s, Действие: %s, Запрос: %s\n", table, action, status)
}
