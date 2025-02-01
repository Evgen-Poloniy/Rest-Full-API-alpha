package main

import (
	"fmt"
)

func input() {
	var chose int

	for {
		clearConsole()
		printActions()
		fmt.Scanln(&chose)
		clearConsole()

		switch chose {
		case 1:
			makeRequest("GET", ipConfig.Ip, ipConfig.Port, "/getCountOfUsers")
			fmt.Println("Нажмите Enter для продолжения...")
			fmt.Scanln()

		case 2:
			makeRequest("GET", ipConfig.Ip, ipConfig.Port, "/getRecordByID?id=1")
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
