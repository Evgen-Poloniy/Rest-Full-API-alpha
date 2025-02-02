package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func makeTransaction(w http.ResponseWriter, r *http.Request) {
	senderID := r.URL.Query().Get("sender_id")
	receiverID := r.URL.Query().Get("receiver_id")
	amountStr := r.URL.Query().Get("amount")

	if senderID == "" || receiverID == "" || amountStr == "" {
		getResponseError(w, http.StatusBadRequest, "Пропущены параметры транзакции")
		return
	}

	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil || amount <= 0 {
		getResponseError(w, http.StatusBadRequest, "Некорретная сумма транзакции")
		return
	}

	tx, err := db.Begin()
	if err != nil {
		getResponseError(w, http.StatusInternalServerError, "Ошибка при запуске транзакции")
		return
	}

	_, err = tx.Exec("UPDATE users SET balance = balance - ? WHERE id = ?", amount, senderID)
	if err != nil {
		tx.Rollback()
		getResponseError(w, http.StatusInternalServerError, "Ошибка при обновлении баланса отправителя")
		return
	}

	_, err = tx.Exec("UPDATE users SET balance = balance + ? WHERE id = ?", amount, receiverID)
	if err != nil {
		tx.Rollback()
		getResponseError(w, http.StatusInternalServerError, "Ошибка при обновлении баланса получателя")
		return
	}

	err = tx.Commit()
	if err != nil {
		getResponseError(w, http.StatusInternalServerError, "Ошибка записи транзакции")
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Printf("Действие: транзакция на сумму %f рублей от %s к %s", amount, senderID, receiverID)
}
