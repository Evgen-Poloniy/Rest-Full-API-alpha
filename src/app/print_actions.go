package main

import (
	"fmt"
)

func printStatusConnection(status bool) string {
	if status {
		return "\033[32mСоединение установлено\033[0m"
	}

	return "\033[31mНет соединения\033[0m"
}

func printActions() {
	fmt.Println("*------------------------------------*")
	fmt.Println("| IP: ", ipConfig.Ip)
	fmt.Println("| PORT: ", ipConfig.Port)
	fmt.Println("| Status: ", printStatusConnection(ipConfig.Status))
	fmt.Println("*------------------------------------*")
	fmt.Println("Выберите действие:")
	if ipConfig.Status {
		fmt.Println("1. Создать новую учетную запись")
		fmt.Println("2. Удалить учетную запись")
		fmt.Println("3. Получить запись по ID")
		fmt.Println("4. Получить все записи")
		fmt.Println("5. Получить количество пользователей")
		fmt.Println("6. Сделать транзакцию")
		fmt.Println("7. Обновить баланс")
		fmt.Println("")
	}
	fmt.Println("c. Проверить соединение")
	fmt.Println("r. Изменить параметры сети")
	fmt.Println("q. Выйти")
}
