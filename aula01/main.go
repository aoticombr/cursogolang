package main

import (
	"fmt"
	"os"
)

const msg_erro = "deu um [%s] %s %d xxxxxxx"

func LerArquivo() (*os.File, error) {
	arq, err := os.Open("rotina1.go")
	if err != nil {
		return nil, err
	}
	return arq, nil
}

func main() {
	fmt.Println()

	_, err := LerArquivo()
	if err != nil {
		fmt.Println(err.Error())
		err = fmt.Errorf(msg_erro, err.Error(), "aconteceu nao funcao ler arquivo", 404)
		fmt.Println(err)
		err = fmt.Errorf(msg_erro, err.Error(), "hhhhh", 500)
		fmt.Println(err)

		panic("deu um erro muito grave, erro:" + err.Error())
	}
	fmt.Println("Arquivo aberto com sucesso")
	//panic("aaaaaaa")
	//arq, err = os.Open("")

}

func init() {
	fmt.Println("Initializing...")
}
