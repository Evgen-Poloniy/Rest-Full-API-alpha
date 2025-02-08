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
		fmt.Println("b. Назад")

		var choseTable string = ""

		fmt.Scanln(&choseTable)
		clearConsole()

		switch choseTable {
		case "1":
			return "users"
		case "2":
			return "transactions"
		case "b":
			return "nil"
		default:
			fmt.Println("Неверный ввод")
			waitInput()
			clearConsole()
		}
	}
}

func choseCurrency() (string, bool) {
	for {
		fmt.Println("Выберите валюту, в которой вы хотите получить ответ:")
		fmt.Println("1. Рубль (RUB)")
		fmt.Println("2. Доллар США (USD)")
		fmt.Println("3. Евро (EUR)")
		fmt.Println("4. Юань (CNY)")
		fmt.Println("5. Иен (JPY)")
		fmt.Println("6. Вон (KRW)")
		fmt.Println("7. Фунт стерлингов (GBP)")
		fmt.Println("8. Индийский рупий(INR)")
		fmt.Println("9. Тенге (KZT)")
		fmt.Println("10. Киргизский сом (KGS)")
		fmt.Println("11. Узбекский сум (UZS)")
		fmt.Println("12. Сомони (TJS)")
		fmt.Println("13. Гривна (UAH)")
		fmt.Println("14. Белорусский рубль(BYN)")
		fmt.Println("15. Туретская лира(TRY)")

		fmt.Println("")
		fmt.Println("b. Назад")

		var choseTable string = ""

		fmt.Scanln(&choseTable)
		clearConsole()

		switch choseTable {
		case "1":
			return "RUB", false
		case "2":
			return "USD", false
		case "3":
			return "EUR", false
		case "4":
			return "CNY", false
		case "5":
			return "JPY", false
		case "6":
			return "KRW", false
		case "7":
			return "GBP", false
		case "8":
			return "INR", false
		case "9":
			return "KZT", false
		case "10":
			return "KGS", false
		case "11":
			return "UZS", false
		case "12":
			return "TJS", false
		case "13":
			return "UAH", false
		case "14":
			return "BYN", false
		case "15":
			return "TRY", false
		case "b":
			return "nil", true
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
				currency, goBack := choseCurrency()
				if goBack {
					continue
				}
				var balance string = inputParametrs("Введите баланс в выбранной валюте:")
				makeRequest("POST", ipConfig.Ip, ipConfig.Port, "/users/createRecordOfUser?balance="+balance+"&currency="+currency)
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
						currency, goBack := choseCurrency()
						if goBack {
							continue
						}
						makeRequest("GET", ipConfig.Ip, ipConfig.Port, "/users/getRecordByID?user_id="+id+"&currency="+currency)
					} else {
						currency, goBack := choseCurrency()
						if goBack {
							continue
						}
						makeRequest("GET", ipConfig.Ip, ipConfig.Port, "/transactions/getRecordByID?transaction_id="+id+"&currency="+currency)
					}
					waitInput()
				}
			case "4":
				var table string = choseTable()
				if table != "nil" {
					var limit string = inputParametrs("Ввведите ограничение на кол-во записей:")
					if table == "users" {
						currency, goBack := choseCurrency()
						if goBack {
							continue
						}
						makeRequest("GET", ipConfig.Ip, ipConfig.Port, "/users/getAllRecords?limit="+limit+"&currency="+currency)
					} else {
						currency, goBack := choseCurrency()
						if goBack {
							continue
						}
						makeRequest("GET", ipConfig.Ip, ipConfig.Port, "/transactions/getAllRecords?limit="+limit+"&currency="+currency)
					}
					waitInput()
				}
			case "5":
				makeRequest("GET", ipConfig.Ip, ipConfig.Port, "/users/getCountOfRecords")
				waitInput()
			case "6":
				currency, goBack := choseCurrency()
				if goBack {
					continue
				}
				var idSender string = inputParametrs("Введите id отправителя:")
				var idReceiver string = inputParametrs("Введите id получателя:")
				var amount string = inputParametrs("Введите сумму транзакции в выбранной валюте:")
				makeRequest("POST", ipConfig.Ip, ipConfig.Port, "/users/makeTransaction?sender_id="+idSender+"&receiver_id="+idReceiver+"&amount="+amount+"&currency="+currency)
				waitInput()
			case "7":
				currency, goBack := choseCurrency()
				if goBack {
					continue
				}
				var id string = inputParametrs("Введите id пользователя:")
				var updateBalance string = inputParametrs("Введите сумму начисления(>0)/списания(<0) в выбранной валюте:")
				makeRequest("POST", ipConfig.Ip, ipConfig.Port, "/users/updateBalanceByID?user_id="+id+"&update_balance="+updateBalance+"&currency="+currency)
				waitInput()
			case "o":
				var password string = inputParametrs("Введите пароль от базы данных:")
				fmt.Println(password)
				openOrCloseDB("openDB", password)
				waitInput()
			case "p":
				var password string = inputParametrs("Введите пароль от базы данных:")
				openOrCloseDB("closeDB", password)
				waitInput()
			case "c":
				ipConfig.Status = checkConnection()
				PrintMessageAboutStatusConnection()
			case "r":
				changeIpConfig()
				ipConfig.Status = checkConnection()
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
				ipConfig.Status = checkConnection()
				PrintMessageAboutStatusConnection()
			case "r":
				changeIpConfig()
				ipConfig.Status = checkConnection()
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
