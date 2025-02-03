package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func checkConnection(ipConfig *IpConfig) {
	var client = &http.Client{}
	resp, err := client.Get("http://" + ipConfig.Ip + ":" + strconv.Itoa(ipConfig.Port) + "/checkConnection")
	if err != nil {
		ipConfig.Status = false
		log.Println(err)
		return
	}
	defer resp.Body.Close()
	ipConfig.Status = resp.StatusCode == http.StatusOK
}

func PrintMessageAboutStatusConnection() {
	if !ipConfig.Status {
		fmt.Println("Нет соединения. Проверьте данные ip-адреса и порта или повторите попытку позже")
		waitInput()
	}
}
