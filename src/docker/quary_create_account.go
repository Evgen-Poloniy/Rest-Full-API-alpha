package main

/*
func createRecord(w http.ResponseWriter, r *http.Request) {
	// Парсим баланс из параметров запроса
	balanceStr := r.URL.Query().Get("balance")
	if balanceStr == "" {
		getResponseError(w, http.StatusBadRequest, "Не указан баланс")
		return
	}

	// Преобразуем строку в число
	balance, err := strconv.ParseFloat(balanceStr, 64)
	if err != nil {
		getResponseError(w, http.StatusBadRequest, "Некорректный формат баланса")
		return
	}

	// Выполняем вставку записи
	result, err := db.Exec("INSERT INTO users (balance) VALUES (?)", balance)
	if err != nil {
		getResponseError(w, http.StatusInternalServerError, "Ошибка при вставке записи")
		return
	}

	// Получаем ID последней вставленной записи
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		responseError(w, http.StatusInternalServerError, "Ошибка при получении ID записи")
		return
	}
}
*/
