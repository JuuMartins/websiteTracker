package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	displatIntro()
	displayMenu()
	command := readCommand()

	switch command {
	case 1:
		startMonitoring()
	case 2:
		displayLogs()
	case 0:
		exitProgram()
	default:
		fmt.Println("Não conheço este comando.")
		os.Exit(-1)
	}
}

func displatIntro() {
	name := "Juliana"
	version := 1.5
	fmt.Println("Olá Sra", name, "!")
	fmt.Println("Programa Versão", version)
}

func displayMenu() {
	fmt.Println("1. INICIAR MONITORAMENTO")
	fmt.Println("2. EXIBIR LOGS")
	fmt.Println("0. SAIR DO PROGRAMA")
}

func readCommand() int {
	var commandRead int
	fmt.Scan(&commandRead)
	fmt.Println("O comando escolhido foi", commandRead)

	return commandRead
}

func startMonitoring() {
	fmt.Println("Monitorando...")
	site := "https://random-status-code.herokuapp.com/"
	response, _ := http.Get(site)

	if response.StatusCode == 200 {
		fmt.Println("Site", site, "carregado com sucesso")
	} else {
		fmt.Println("O site", site, "está com problema. Status Code:",
			response.StatusCode)

	}
}

func displayLogs() {
	fmt.Println("Exibindo logs...")
}

func exitProgram() {
	fmt.Println("Saindo do programa")
	os.Exit(0)
}
