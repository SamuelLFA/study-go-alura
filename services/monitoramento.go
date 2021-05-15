package services

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const tentativas = 3
const delay = 5

func IniciaMonitoramento() {
	fmt.Println("Iniciando monitoramento...")
	fmt.Println()
	sites := leSitesdeArquivo()

	for i := 0; i < tentativas; i++ {
		for _, site := range sites {
			testaSite(site)
		}
		fmt.Println()
		time.Sleep(delay * time.Second)
	}

}

func testaSite(site string) {
	fmt.Println("Testando site:", site, "....")
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro na requisição:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "carregado com sucesso!")
		registraLog(site, true)
	} else {
		fmt.Println("Site:", site, "esta com problemas. Status code:", resp.StatusCode)
		registraLog(site, false)
	}
}

func leSitesdeArquivo() []string {
	slice := make([]string, 0)
	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
		return slice
	}

	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		slice = append(slice, linha)

		if err == io.EOF {
			break
		}
	}

	arquivo.Close()
	return slice
}

func registraLog(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Erro ao abrir arquivo:", err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}
