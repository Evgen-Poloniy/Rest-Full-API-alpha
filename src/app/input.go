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
	var chose string

	for {
		clearConsole()
		printActions()
		fmt.Scanln(&chose)
		clearConsole()

		if ipConfig.Status {
			if chose == "1" {
				var balance string = inputParametrs("Введите баланс:")
				makeRequest("POST", ipConfig.Ip, ipConfig.Port, "/users/createRecord?balance="+balance)
				fmt.Println("Нажмите Enter для продолжения...")
				fmt.Scanln()
			} else if chose == "2" {
				var id string = inputParametrs("Введите id пользователя:")
				makeRequest("DELETE", ipConfig.Ip, ipConfig.Port, "/users/deleteRecordById?id="+id)
				fmt.Println("Запись с id =", id, "была удалена")
				fmt.Println("Нажмите Enter для продолжения...")
				fmt.Scanln()
			} else if chose == "3" {
				var id string = inputParametrs("Введите id пользователя:")
				makeRequest("GET", ipConfig.Ip, ipConfig.Port, "/users/getRecordByID?id="+id)
				fmt.Println("Нажмите Enter для продолжения...")
				fmt.Scanln()
			} else if chose == "4" {
				var limit string = inputParametrs("Ввведите ограничение на кол-во записей:")
				makeRequest("GET", ipConfig.Ip, ipConfig.Port, "/users/getAllRecords?limit="+limit)
				fmt.Println("Нажмите Enter для продолжения...")
				fmt.Scanln()
			} else if chose == "6" {
				makeRequest("GET", ipConfig.Ip, ipConfig.Port, "/users/getCountOfRecords")
				fmt.Println("Нажмите Enter для продолжения...")
				fmt.Scanln()
			} else if chose == "7" {
				var idSender string = inputParametrs("Введите id отправителя:")
				var idReceiver string = inputParametrs("Введите id получателя:")
				var amount string = inputParametrs("Введите сумму транзакции:")
				makeRequest("POST", ipConfig.Ip, ipConfig.Port, "/users/makeTransaction?sender_id="+idSender+"&receiver_id="+idReceiver+"&amount="+amount)
				fmt.Println("Произведен перевод средств на сумму", amount, "от пользователя с id =", idSender, "к пользователю с id =", idReceiver)
				fmt.Println("Нажмите Enter для продолжения...")
				fmt.Scanln()
			}
		}
		if chose == "9" {
			checkConnection(&ipConfig)
		} else if chose == "10" {
			changeIpConfig()
			checkConnection(&ipConfig)
		} else if chose == "q" {
			fmt.Println("Произведен выход...")
			return
		} else {
			fmt.Println("Неверный ввод")
			fmt.Println("Нажмите Enter для продолжения...")
			fmt.Scanln()
		}
	}
}
