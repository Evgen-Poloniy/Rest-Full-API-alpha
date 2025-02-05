package main

import (
	"fmt"
)

func waitInput() {
	fmt.Println("Нажмите Enter для продолжения...")
	fmt.Scanln()
}

func inputParametrs(message string) string {
	clearConsole()
	fmt.Println(message)
	var parametrs string
	fmt.Scanln(&parametrs)
	clearConsole()
	return parametrs
}

func choseTable() string {
	for {
		fmt.Println("Выберите таблицу:")
		fmt.Println("1. Пользователи")
		fmt.Println("2. Транзакции")
		fmt.Println("")
		fmt.Println("3. Назад")

		var choseTable int = 3

		fmt.Scanln(&choseTable)
		clearConsole()

		switch choseTable {
		case 1:
			return "users"
		case 2:
			return "transactions"
		case 3:
			return "nil"
		default:
			fmt.Println("Неверный ввод")
			waitInput()
			clearConsole()
		}
	}
}

func input() {
	var chose string

	for {
		clearConsole()
		printActions()
		fmt.Scanln(&chose)
		clearConsole()

		if ipConfig.Status {
			switch chose {
			case "1":
				var balance string = inputParametrs("Введите баланс:")
				makeRequest("POST", ipConfig.Ip, ipConfig.Port, "/users/createRecordOfUser?balance="+balance)
				waitInput()
			case "2":
				var id string = inputParametrs("Введите id пользователя:")
				makeRequest("DELETE", ipConfig.Ip, ipConfig.Port, "/users/deleteRecordByID?user_id="+id)
				waitInput()
			case "3":
				var table string = choseTable()
				if table != "nil" {
					var id string = inputParametrs("Введите id:")
					if table == "users" {
						makeRequest("GET", ipConfig.Ip, ipConfig.Port, "/users/getRecordByID?user_id="+id)
					} else {
						makeRequest("GET", ipConfig.Ip, ipConfig.Port, "/transactions/getRecordByID?transaction_id="+id)
					}
					waitInput()
				}
			case "4":
				var table string = choseTable()
				if table != "nil" {
					var limit string = inputParametrs("Ввведите ограничение на кол-во записей:")
					if table == "users" {
						makeRequest("GET", ipConfig.Ip, ipConfig.Port, "/users/getAllRecords?limit="+limit)
					} else {
						makeRequest("GET", ipConfig.Ip, ipConfig.Port, "/transactions/getAllRecords?limit="+limit)
					}
					waitInput()
				}
			case "5":
				makeRequest("GET", ipConfig.Ip, ipConfig.Port, "/users/getCountOfRecords")
				waitInput()
			case "6":
				var idSender string = inputParametrs("Введите id отправителя:")
				var idReceiver string = inputParametrs("Введите id получателя:")
				var amount string = inputParametrs("Введите сумму транзакции:")
				makeRequest("POST", ipConfig.Ip, ipConfig.Port, "/users/makeTransaction?sender_id="+idSender+"&receiver_id="+idReceiver+"&amount="+amount)
				waitInput()
			case "7":
				var id string = inputParametrs("Введите id пользователя:")
				var updateBalance string = inputParametrs("Введите сумму начисления(>0)/списания(<0):")
				makeRequest("POST", ipConfig.Ip, ipConfig.Port, "/users/updateBalanceByID?user_id="+id+"&update_balance="+updateBalance)
				waitInput()
			case "c":
				checkConnection(&ipConfig)
				PrintMessageAboutStatusConnection()
			case "r":
				changeIpConfig()
				checkConnection(&ipConfig)
			case "q":
				fmt.Println("Произведен выход...")
				return
			default:
				fmt.Println("Неверный ввод")
				waitInput()
			}
		} else {
			switch chose {
			case "c":
				checkConnection(&ipConfig)
				PrintMessageAboutStatusConnection()
			case "r":
				changeIpConfig()
				checkConnection(&ipConfig)
			case "q":
				fmt.Println("Произведен выход...")
				return
			default:
				fmt.Println("Неверный ввод")
				waitInput()
			}
		}

	}
}
