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

	var mainTable string = parts[1]
	var additionalTable string = ""
	var action string = parts[2]
	var statusMainQuary bool = false
	var statusAdditionalQuary bool = false
	var additionalAction string = ""

	switch action {
	case "createRecordOfUser":
		statusMainQuary, statusAdditionalQuary = createRecordOfUsers(w, r, mainTable)
		additionalTable = "users"
		additionalAction = "getRecordByID"
	case "deleteRecordByID":
		statusMainQuary = deleteRecordById(w, r, mainTable)
	case "getRecordByID":
		statusMainQuary = getRecordByID(w, r, mainTable)
	case "getAllRecords":
		statusMainQuary = getAllRecords(w, r, mainTable)
	case "getCountOfRecords":
		statusMainQuary = getCountOfRecords(w, mainTable)
	case "makeTransaction":
		statusMainQuary, statusAdditionalQuary = makeTransaction(w, r, mainTable)
		additionalTable = "transactions"
		additionalAction = "getRecordByID"
	case "updateBalanceByID":
		statusMainQuary, statusAdditionalQuary = updateBalanceByID(w, r, mainTable)
		additionalTable = "transactions"
		additionalAction = "getRecordByID"
	default:
		getResponseError(w, http.StatusNotFound, "Неизвестное действие")
	}

	var status string
	if statusMainQuary {
		status = "Успешный"
		if statusAdditionalQuary {
			defer log.Printf("Таблица: %s, Действие: %s, Запрос: %s\n", additionalTable, additionalAction, status)
		}
	} else {
		status = "Неудачный"
	}
	log.Printf("Таблица: %s, Действие: %s, Запрос: %s\n", mainTable, action, status)
}
