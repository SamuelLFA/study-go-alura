package services

import (
	"fmt"
	"io/ioutil"
)

func Imprime() {

	arquivo, err := ioutil.ReadFile("log.txt")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(arquivo))
}
