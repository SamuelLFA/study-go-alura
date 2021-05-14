package services

import (
	"fmt"
	"net/http"
)

func IniciaMonitoramento() {
	fmt.Println("Iniciando monitoramento...")
	site := "https://www.alura.com.br"
	resp, _ := http.Get(site)

	fmt.Println(resp.StatusCode)
}
