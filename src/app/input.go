package main

import (
	"fmt"
	"strconv"
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

type Tables map[string]string

var tables = Tables{
	"1": "users",
	"2": "transactions",
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

		if choseTable == "b" {
			return "nil"
		} else if _, has := tables[choseTable]; has {
			return tables[choseTable]
		} else {
			fmt.Println("Неверный ввод")
			waitInput()
			clearConsole()
		}
	}
}

type CurrencySet map[string][2]string

var currencies = CurrencySet{
	"1":  {"RUB", "Рубль"},
	"2":  {"USD", "Доллар США"},
	"3":  {"EUR", "Евро"},
	"4":  {"CNY", "Юань"},
	"5":  {"JPY", "Иен"},
	"6":  {"KRW", "Вон"},
	"7":  {"GBP", "Фунт стерлингов"},
	"8":  {"INR", "Индийский рупий"},
	"9":  {"KZT", "Тенге"},
	"10": {"KGS", "Киргизский сом"},
	"11": {"UZS", "Узбекский сум"},
	"12": {"TJS", "Сомони"},
	"13": {"UAH", "Гривна"},
	"14": {"BYN", "Белорусский рубль"},
	"15": {"TRY", "Туретская лира"},
}

func printCurrency() {
	var lenth int = len(currencies) + 1
	for i := 1; i < lenth; i++ {
		var iStr string = strconv.Itoa(i)
		fmt.Println(iStr+".", currencies[iStr][1], "("+currencies[iStr][0]+")")
	}
}

func choseCurrency() (string, bool) {
	for {
		fmt.Println("Выберите валюту:")
		printCurrency()

		fmt.Println("")
		fmt.Println("b. Назад")

		var choseTable string = ""

		fmt.Scanln(&choseTable)
		clearConsole()

		if choseTable == "b" {
			return "nil", true
		} else if _, has := currencies[choseTable]; has {
			return currencies[choseTable][0], false
		} else {
			fmt.Println("Неверный ввод")
			waitInput()
			clearConsole()
		}
	}
}

func choseSortingParametrs() (string, bool) {
	for {
		fmt.Println("Выберите парамерт, для сортировки:")
		fmt.Println("1. id пользователя/транзакции")
		fmt.Println("2. Баланс/Сумма транзакции")
		fmt.Println("3. Время транзакции")
		fmt.Println("")
		fmt.Println("b. Назад")

		var sortingParametrs string = ""

		fmt.Scanln(&sortingParametrs)
		clearConsole()

		switch sortingParametrs {
		case "1":
			return "id", false
		case "2":
			return "amount", false
		case "3":
			return "transaction_time", false
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
					parametr, goBack := choseSortingParametrs()
					if goBack {
						continue
					}
					var order string = inputParametrs("Сортировать по возрастанию? [y/n]:")
					if order != "y" || order == "" {
						order = "ASC"
					} else {
						order = "DESC"
					}
					var limit string = inputParametrs("Ввведите ограничение на кол-во записей или оставьте пустым:")
					var offset string = inputParametrs("Ввведите отступ записей (пагинацию) или оставьте пустым:")
					if table == "users" {
						currency, goBack := choseCurrency()
						if goBack {
							continue
						}
						makeRequest("GET", ipConfig.Ip, ipConfig.Port, "/users/getAllRecords?currency="+currency+"&parametr="+parametr+"&order="+order+"&limit="+limit+"&offset="+offset)
					} else {
						currency, goBack := choseCurrency()
						if goBack {
							continue
						}
						makeRequest("GET", ipConfig.Ip, ipConfig.Port, "/transactions/getAllRecords?currency="+currency+"&parametr="+parametr+"&order="+order+"&limit="+limit+"&offset="+offset)
					}
					waitInput()
				}
			case "5":
				makeRequest("GET", ipConfig.Ip, ipConfig.Port, "/users/getCountOfRecords")
				waitInput()
			case "t":
				currency, goBack := choseCurrency()
				if goBack {
					continue
				}
				var idSender string = inputParametrs("Введите id отправителя:")
				var idReceiver string = inputParametrs("Введите id получателя:")
				var amount string = inputParametrs("Введите сумму транзакции в выбранной валюте:")
				makeRequest("POST", ipConfig.Ip, ipConfig.Port, "/users/makeTransaction?sender_id="+idSender+"&receiver_id="+idReceiver+"&amount="+amount+"&currency="+currency)
				waitInput()
			case "u":
				currency, goBack := choseCurrency()
				if goBack {
					continue
				}
				var id string = inputParametrs("Введите id пользователя:")
				var updateBalance string = inputParametrs("Введите сумму начисления(>0)/списания(<0) в выбранной валюте:")
				makeRequest("POST", ipConfig.Ip, ipConfig.Port, "/users/updateBalanceByID?user_id="+id+"&update_balance="+updateBalance+"&currency="+currency)
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
