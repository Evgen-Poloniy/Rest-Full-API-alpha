package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func openOrCloseDB(action string, password string) {
	if action == "openDB" {
		fmt.Println("Подождите, пока открывается база данных")
	}
	var client = &http.Client{}
	resp, err := client.Get("http://" + ipConfig.Ip + ":" + strconv.Itoa(ipConfig.Port) + "/" + action + "?password=" + password)
	if err != nil {
		fmt.Println("База данных не доступна")
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		if action == "openDB" {
			fmt.Println("База данных открыта")
		} else {
			fmt.Println("База данных закрыта")
		}
	} else if resp.StatusCode == http.StatusUnauthorized {
		fmt.Println("Неверный пароль")
	} else if resp.StatusCode == http.StatusBadRequest {
		fmt.Println("Неправильный запрос")
		fmt.Println("http://" + ipConfig.Ip + ":" + strconv.Itoa(ipConfig.Port) + "/" + action + "?password=" + password)
	}
}
