package main

import (
	"net/http"
	"strconv"
)

func makeTransaction(w http.ResponseWriter, r *http.Request, table string) bool {
	senderID := r.URL.Query().Get("sender_id")
	receiverID := r.URL.Query().Get("receiver_id")
	amountStr := r.URL.Query().Get("amount")

	if senderID == "" || receiverID == "" || amountStr == "" {
		getResponseError(w, http.StatusBadRequest, "Пропущены параметры транзакции")
		return false
	}

	if senderID == receiverID {
		getResponseError(w, http.StatusBadRequest, "Указан один и тот же user_id")
		return false
	}

	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil || amount <= 0 {
		getResponseError(w, http.StatusBadRequest, "Некорретная сумма транзакции")
		return false
	}

	reqSenderID, _ := http.NewRequest("GET", "/users/getRecordByID?user_id="+senderID, nil)
	reqReceiverID, _ := http.NewRequest("GET", "/users/getRecordByID?user_id="+receiverID, nil)

	_, existSenderID := ckeckRecordOfUsersByID(w, reqSenderID, "users")
	_, existReceiverID := ckeckRecordOfUsersByID(w, reqReceiverID, "users")

	if !existSenderID || !existReceiverID {
		if !existSenderID {
			getResponseError(w, http.StatusNotFound, "Отправитель не найден")
		}
		if !existReceiverID {
			getResponseError(w, http.StatusNotFound, "Получатель не найден")
		}
		return false
	}

	var valueOfBalance float64 = 0
	errQuary := db.QueryRow("SELECT balance FROM "+table+" WHERE user_id = ?", senderID).Scan(&valueOfBalance)
	if errQuary != nil {
		getResponseError(w, http.StatusInternalServerError, "Ошибка при запросе баланса пользователя")
		return false
	}

	if valueOfBalance < amount {
		getResponseError(w, http.StatusBadRequest, "Не хратает средств на счете")
		return false
	}

	tx, err := db.Begin()
	if err != nil {
		getResponseError(w, http.StatusInternalServerError, "Ошибка при запуске транзакции")
		return false
	}

	_, err = tx.Exec("UPDATE "+table+" SET balance = balance - ?, transaction_time = CURRENT_TIMESTAMP WHERE user_id = ?", amount, senderID)
	if err != nil {
		tx.Rollback()
		getResponseError(w, http.StatusInternalServerError, "Ошибка при обновлении баланса отправителя")
		return false
	}

	_, err = tx.Exec("UPDATE "+table+" SET balance = balance + ?, transaction_time = CURRENT_TIMESTAMP WHERE user_id = ?", amount, receiverID)
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

	transactionSenderID := TransactionParameters{
		ID:     senderID,
		Type:   "transfer_in",
		Amount: amountStr,
	}

	createRecordOfTransaction(w, "transactions", &transactionSenderID)

	transactionReceiverID := TransactionParameters{
		ID:     receiverID,
		Type:   "transfer_out",
		Amount: amountStr,
	}

	createRecordOfTransaction(w, "transactions", &transactionReceiverID)

	return true
}
