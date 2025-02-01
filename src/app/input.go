package main

import (
	"fmt"
)

func inputParametrs(message string) string {
	clearConsole()
	fmt.Println(message)
	var parametrs string
	fmt.Scanln(&parametrs)
	clearConsole()
	return parametrs
}

func input() {
	var chose int

	for {
		clearConsole()
		printActions()
		fmt.Scanln(&chose)
		clearConsole()

		switch chose {
		case 1:
			var balance string = inputParametrs("Введите баланс")

			makeRequest("POST", ipConfig.Ip, ipConfig.Port, "/createRecord?balance="+balance)
			fmt.Println("Нажмите Enter для продолжения...")
			fmt.Scanln()

		case 2:
			var id string = inputParametrs("Введите id пользователя")

			makeRequest("GET", ipConfig.Ip, ipConfig.Port, "/getRecordByID?id="+id)
			fmt.Println("Нажмите Enter для продолжения...")
			fmt.Scanln()

		case 3:
			var limit string = inputParametrs("Ввведите ограничение на кол-во записей")
			makeRequest("GET", ipConfig.Ip, ipConfig.Port, "/getAllRecords?limit="+limit)
			fmt.Println("Нажмите Enter для продолжения...")
			fmt.Scanln()

		case 6:
			makeRequest("GET", ipConfig.Ip, ipConfig.Port, "/getCountOfUsers")
			fmt.Println("Нажмите Enter для продолжения...")
			fmt.Scanln()

		case 7:
			changeIpConfig()

		case 8:
			fmt.Println("Произведен выход...")
			return

		default:
			fmt.Println("Неверный ввод")
			fmt.Println("Нажмите Enter для продолжения...")
			fmt.Scanln()
		}
	}
}
