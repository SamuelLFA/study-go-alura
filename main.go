package main

import (
	"fmt"
	"os"

	"github.com/SamuelLFA/study-go-alura/services"
)

func main() {

	for {
		exibeMenu()
		comando := leComando()

		switch comando {
		case 1:
			services.IniciaMonitoramento()
		case 2:
			services.Imprime()
		case 0:
			fmt.Println("Fechando...")
			os.Exit(0)
		default:
			fmt.Println("Fechando...")
			os.Exit(0)
		}
	}
}

func exibeMenu() {
	fmt.Println("1- Iniciar monitoramento")
	fmt.Println("2- Exibir logs")
	fmt.Println("0- Sair do programa")
	fmt.Println()
}

func leComando() int {
	var comando int

	fmt.Scan(&comando)
	fmt.Println()

	return comando
}
