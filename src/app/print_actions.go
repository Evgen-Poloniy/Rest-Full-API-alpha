package main

import (
	"fmt"
)

func printActions() {
	fmt.Println("*----------------------*")
	fmt.Println("| IP: ", ipConfig.Ip)
	fmt.Println("| PORT: ", ipConfig.Port)
	fmt.Println("*----------------------*")
	fmt.Println("Выберите действие:")
	fmt.Println("1. Создать новую учетную запись")
	fmt.Println("2. Получить запись по ID")
	fmt.Println("3. Получить все записи")
	fmt.Println("6. Получить количество пользователей")
	fmt.Println("7. Изменить параметры сети")

	fmt.Println("8. Выйти")
}
