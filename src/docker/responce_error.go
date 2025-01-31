package main

import (
	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

// Структура для вывода ошибки в формате JSON
type ErrorResponse struct {
	Error   int    `json:"error"`
	Message string `json:"message"`
}

// Вывод кода ошибки в формате JSON
func responseError(w http.ResponseWriter, statusCode int, errMsg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := ErrorResponse{
		Error:   statusCode,
		Message: errMsg,
	}

	responseJSON, _ := json.MarshalIndent(response, "", "  ")

	w.Write(append(responseJSON, '\n'))
}
