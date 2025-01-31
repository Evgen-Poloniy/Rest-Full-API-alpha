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
			makeRequest("GET", "localhost", 8080, "/getRecordByID?id=1")
			fmt.Println("Нажмите Enter для продолжения...")
			fmt.Scanln()

		case 2:
			makeRequest("GET", "localhost", 8080, "/getCountOfUsers")
			fmt.Println("Нажмите Enter для продолжения...")
			fmt.Scanln()

		case 7:
			fmt.Println("Произведен выход...")
			return

		default:
			fmt.Println("Неверный ввод")
			fmt.Println("Нажмите Enter для продолжения...")
			fmt.Scanln()
		}
	}
}
