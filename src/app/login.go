package main

import (
	"fmt"
	"os"
	"strconv"
)

type IpConfig struct {
	Ip     string
	Port   int
	Status bool
}

var ipConfig IpConfig

func fileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)
}

func changeIpConfig() {
	fmt.Println("Введите ip-адресс сервера:")
	var inputedIp string
	fmt.Scanln(&inputedIp)

	fmt.Println("Введите порт:")
	var inputedPort int
	fmt.Scanln(&inputedPort)

	if inputedIp == "" && inputedPort == 0 {
		return
	}

	fmt.Println("Принять изменения? y/n:")
	var answer string
	fmt.Scanln(&answer)

	if answer == "y" {
		file, err := os.Create("./data/ipconfig.txt")
		if err != nil {
			fmt.Println("Ошибка при создании файла", err)
			return
		}

		defer file.Close()

		if inputedIp != "" {
			ipConfig.Ip = inputedIp
		}
		if inputedPort != 0 {
			ipConfig.Port = inputedPort
		}

		file.WriteString(ipConfig.Ip + "\n")
		file.WriteString(strconv.Itoa(ipConfig.Port) + "\n")
	} else {
		return
	}
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
	ipConfig.Status = checkConnection()
}
