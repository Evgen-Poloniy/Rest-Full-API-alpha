package main

import (
	"net/http"
	"strconv"
)

func updateBalanceByID(w http.ResponseWriter, r *http.Request, table string) bool {
	userIDStr := r.URL.Query().Get("user_id")
	updateBalanceStr := r.URL.Query().Get("update_balance")

	if userIDStr == "" {
		getResponseError(w, http.StatusBadRequest, "Параметр user_id обязателен")
		return false
	}

	if updateBalanceStr == "" {
		getResponseError(w, http.StatusBadRequest, "Параметр update_balance обязателен")
		return false
	}

	updateBalance, err := strconv.ParseFloat(updateBalanceStr, 64)
	if err != nil || updateBalance == 0 {
		getResponseError(w, http.StatusBadRequest, "Некорретная сумма транзакции")
		return false
	}

	reqID, _ := http.NewRequest("GET", "/users/getRecordByID?user_id="+userIDStr, nil)
	_, exist := ckeckRecordOfUsersByID(w, reqID, table)

	if !exist {
		getResponseError(w, http.StatusNotFound, "Пользователь не найден")
		return false
	}

	tx, err := db.Begin()
	if err != nil {
		getResponseError(w, http.StatusInternalServerError, "Ошибка при запуске транзакции")
		return false
	}

	var typeOperation string

	if updateBalance > 0 {
		typeOperation = "deposit"
		_, err = tx.Exec("UPDATE users SET balance = balance + ?, transaction_time = CURRENT_TIMESTAMP WHERE user_id = ?", updateBalance, userIDStr)
		if err != nil {
			tx.Rollback()
			getResponseError(w, http.StatusInternalServerError, "Ошибка при обновлении баланса отправителя")
			return false
		}
	} else {
		typeOperation = "withdraw"
		_, err = tx.Exec("UPDATE users SET balance = balance - ?, transaction_time = CURRENT_TIMESTAMP WHERE user_id = ? AND balance >= ?", updateBalance, userIDStr, updateBalance)
		if err != nil {
			tx.Rollback()
			getResponseError(w, http.StatusInternalServerError, "Ошибка при обновлении баланса отправителя")
			return false
		}
	}

	err = tx.Commit()
	if err != nil {
		getResponseError(w, http.StatusInternalServerError, "Ошибка записи транзакции")
		return false
	}

	transactionSenderID := TransactionParameters{
		ID:     userIDStr,
		Type:   typeOperation,
		Amount: updateBalanceStr,
	}

	createRecordOfTransaction(w, "transactions", &transactionSenderID)

	return true

}
