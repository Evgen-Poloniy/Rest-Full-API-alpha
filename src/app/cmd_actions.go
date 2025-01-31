package main

import (
	"os"
	"os/exec"
	"runtime"
	"strconv"
)

func makeRequest(typeOfRequest string, ip string, port int, request string) {
	var cmd *exec.Cmd
	var url string = "http://" + ip + ":" + strconv.Itoa(port) + request

	cmd = exec.Command("curl", "-X", typeOfRequest, url)

	cmd.Stdout = os.Stdout
	cmd.Run()
}

func clearConsole() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}
