package main

import (
	"net/http"
	"strconv"
)

func makeTransaction(w http.ResponseWriter, r *http.Request) bool {
	senderID := r.URL.Query().Get("sender_id")
	receiverID := r.URL.Query().Get("receiver_id")
	amountStr := r.URL.Query().Get("amount")

	if senderID == "" || receiverID == "" || amountStr == "" {
		getResponseError(w, http.StatusBadRequest, "Пропущены параметры транзакции")
		return false
	}

	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil || amount <= 0 {
		getResponseError(w, http.StatusBadRequest, "Некорретная сумма транзакции")
		return false
	}

	reqSenderID, _ := http.NewRequest("GET", "/users/getRecordByID?user_id="+senderID, nil)
	reqReceiverID, _ := http.NewRequest("GET", "/users/getRecordByID?user_id="+receiverID, nil)

	_, existSenderID := ckeckRecordByID(w, reqSenderID, "users")
	_, existreceiverID := ckeckRecordByID(w, reqReceiverID, "users")

	if !existSenderID && !existreceiverID {
		return false
	}
	tx, err := db.Begin()
	if err != nil {
		getResponseError(w, http.StatusInternalServerError, "Ошибка при запуске транзакции")
		return false
	}

	_, err = tx.Exec("UPDATE users SET balance = balance - ? WHERE user_id = ?", amount, senderID)
	if err != nil {
		tx.Rollback()
		getResponseError(w, http.StatusInternalServerError, "Ошибка при обновлении баланса отправителя")
		return false
	}

	_, err = tx.Exec("UPDATE users SET balance = balance + ? WHERE user_id = ?", amount, receiverID)
	if err != nil {
		tx.Rollback()
		getResponseError(w, http.StatusInternalServerError, "Ошибка при обновлении баланса получателя")
		return false
	}

	err = tx.Commit()
	if err != nil {
		getResponseError(w, http.StatusInternalServerError, "Ошибка записи транзакции")
		return false
	}

	return true
}
