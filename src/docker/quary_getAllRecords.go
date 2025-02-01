package main

import (
	"database/sql"
	"net/http"
)

func getAllRecords(w http.ResponseWriter, r *http.Request) {
	limitStr := r.URL.Query().Get("limit")
	var rows *sql.Rows
	var err error

	if limitStr != "" {
		if limitStr == "1" {
			var record Record
			err := db.QueryRow("SELECT * FROM users ORDER BY id LIMIT 1").Scan(&record.ID, &record.Balance, &record.Time)

			if err != nil {
				getResponseError(w, 404, "Пользователь не найден")
				return
			}

			getResponseRecord(w, record)
			return
		}
		rows, err = db.Query("SELECT * FROM users ORDER BY id LIMIT ?", limitStr)
	} else {
		rows, err = db.Query("SELECT * FROM users ORDER BY id")
	}

	if err != nil {
		getResponseError(w, http.StatusInternalServerError, "Ошибка запроса к БД")
		return
	}
	defer rows.Close()

	var records []Record

	for rows.Next() {
		var record Record
		if err := rows.Scan(&record.ID, &record.Balance, &record.Time); err != nil {
			getResponseError(w, http.StatusInternalServerError, "Ошибка чтения данных")
			return
		}
		records = append(records, record)
	}

	if len(records) == 0 {
		http.Error(w, "Нет данных в таблице", http.StatusNotFound)
		getResponseError(w, http.StatusNotFound, "В таблице нет данных")
		return
	}

	getResponseRecords(w, records)
}
