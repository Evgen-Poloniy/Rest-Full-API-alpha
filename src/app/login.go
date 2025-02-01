package main

import (
	"fmt"
	"os"
	"strconv"
)

type IpConfig struct {
	Ip   string
	Port int
}

var ipConfig IpConfig

func fileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)
}

func changeIpConfig() {
	fmt.Println("Введите ip-адресс сервера:")
	fmt.Scanln(&ipConfig.Ip)

	fmt.Println("Введите порт:")
	fmt.Scanln(&ipConfig.Port)

	file, err := os.Create("./data/ipconfig.txt")
	if err != nil {
		fmt.Println("Ошибка при создании файлы", err)
	}

	defer file.Close()

	file.WriteString(ipConfig.Ip + "\n")
	file.WriteString(strconv.Itoa(ipConfig.Port) + "\n")
}

func logIn() {
	if fileExists("./data/ipconfig.txt") {
		file, err := os.Open("./data/ipconfig.txt")
		if err != nil {
			fmt.Println("Ошибка при открытии файла:", err)
			return
		}
		defer file.Close()

		_, err = fmt.Fscanf(file, "%s\n%d", &ipConfig.Ip, &ipConfig.Port)
		if err != nil {
			fmt.Println("Ошибка при чтении файла:", err)
			return
		}
	} else {
		changeIpConfig()
	}
}
