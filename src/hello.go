package main

import (
	"fmt"
	"os"

	"github.com/SamuelLFA/study-go-alura/services"
)

func main() {
	exibeMenu()
	comando := leComando()

	switch comando {
	case 1:
		services.IniciaMonitoramento()
	case 2:
		fmt.Println("Exibindo logs...")
	case 0:
	default:
		fmt.Println("Fechando...")
		os.Exit(0)
	}
}

func exibeMenu() {
	fmt.Println("1- Iniciar monitoramento")
	fmt.Println("2- Exibir logs")
	fmt.Println("0- Sair do programa")
}

func leComando() int {
	var comando int

	fmt.Scan(&comando)

	return comando
}
