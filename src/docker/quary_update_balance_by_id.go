package main

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
)

func updateBalanceByID(w http.ResponseWriter, r *http.Request, table string) (bool, bool) {
	var userIDStr string = r.URL.Query().Get("user_id")
	var updateBalanceStr string = r.URL.Query().Get("update_balance")
	var currency string = r.URL.Query().Get("currency")

	if userIDStr == "" {
		getResponseError(w, http.StatusBadRequest, "Параметр user_id обязателен")
		return false, false
	}

	if updateBalanceStr == "" {
		getResponseError(w, http.StatusBadRequest, "Параметр update_balance обязателен")
		return false, false
	}

	amount, err := strconv.ParseFloat(updateBalanceStr, 64)
	if err != nil || amount == 0 {
		getResponseError(w, http.StatusBadRequest, "Некорретная сумма транзакции")
		return false, false
	}

	if !updateExchangeRates(w, &amount, &currency) {
		return false, false
	}

	reqID, _ := http.NewRequest("GET", "/users/getRecordByID?user_id="+userIDStr, nil)
	_, exist := ckeckRecordOfUsersByID(w, reqID, table)

	if !exist {
		getResponseError(w, http.StatusNotFound, "Пользователь не найден")
		return false, false
	}

	var valueOfBalance float64 = 0
	errQuary := db.QueryRow("SELECT balance FROM "+table+" WHERE user_id = ?", userIDStr).Scan(&valueOfBalance)
	fmt.Println(valueOfBalance)
	if errQuary != nil {
		getResponseError(w, http.StatusInternalServerError, "Ошибка при запросе баланса пользователя")
		return false, false
	}

	if amount < 0 {
		if valueOfBalance < math.Abs(amount) {
			getResponseError(w, http.StatusBadRequest, "Не хватает средств на счете")
			return false, false
		}
	}

	tx, err := db.Begin()
	if err != nil {
		getResponseError(w, http.StatusInternalServerError, "Ошибка при запуске транзакции")
		return false, false
	}

	var typeOperation string

	_, err = tx.Exec("UPDATE users SET balance = balance + ?, transaction_time = CURRENT_TIMESTAMP WHERE user_id = ?", amount, userIDStr)
	if amount > 0 {
		typeOperation = "deposit"
		if err != nil {
			tx.Rollback()
			getResponseError(w, http.StatusInternalServerError, "Ошибка при обновлении баланса отправителя")
			return false, false
		}
	} else {
		typeOperation = "withdraw"
		if err != nil {
			tx.Rollback()
			getResponseError(w, http.StatusInternalServerError, "Ошибка при обновлении баланса отправителя")
			return false, false
		}
	}

	err = tx.Commit()
	if err != nil {
		getResponseError(w, http.StatusInternalServerError, "Ошибка записи транзакции")
		return false, false
	}

	transactionSenderID := TransactionParameters{
		ID:     userIDStr,
		Type:   typeOperation,
		Amount: updateBalanceStr,
	}

	return true, createRecordOfTransaction(w, "transactions", &transactionSenderID, &currency)
}
