package main

import (
	"fmt"
	"net/http"
	"strings"
)

func dynamicHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")

	if len(parts) == 2 && parts[1] == "checkConnection" {
		checkConnectionHandler(w, r)
		fmt.Printf("Действие: checkConnection")
		return
	}

	if len(parts) < 3 {
		getResponseError(w, http.StatusBadRequest, "Неверный запрос")
		return
	}

	table := parts[1]
	action := parts[2]

	switch action {
	case "createRecord":
		createRecord(w, r, table)
	case "deleteRecordById":
		deleteRecordById(w, r, table)
	case "getRecordByID":
		getRecordByID(w, r, table)
	case "getAllRecords":
		getAllRecords(w, r, table)
	case "getCountOfRecords":
		getCountOfRecords(w, r, table)
	default:
		getResponseError(w, http.StatusNotFound, "Неизвестное действие")
	}

	fmt.Printf("Таблица: %s, Действие: %s\n", table, action)
}
