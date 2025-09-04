package main

import (
	"fmt"
	"os"

	"github.com/aoticombr/golang/component"
)

var teste2 = "joao"

func Soma(a, b int) int {
	return a + b
}
func Dividir(a, b int) float32 {
	if b == 0 {
		return 0
	}
	return float32(a) / float32(b)
}

// LerArquivo
//
//	essa funcao ler arquivos que voce indicar
func LerArquivo(nome string) (arq *os.File, err error) {
	fmt.Println(teste2)
	arq, err = os.Open(nome)
	if err != nil {
		teste5 := "rubens"

		fmt.Println(teste5)
		return nil, err
	}
	//fmt.Println(teste5)
	//fmt.Println(teste3)
	return
}

func main() {
	var texto component.Strings
	texto.Add("asdasdasd")
	texto.Add("sdfdfdf")

	fmt.Println(texto.Text())

	teste3 := "kamila"
	fmt.Println(teste3)
	fmt.Println(teste2)
	fmt.Println("Hello, World!")
	fmt.Println(Soma(5, 10))
	fmt.Println(Dividir(45, 0) + Dividir(10, 2))

	var teste string

	teste = " paulo"

	defer func(x string) {
		fmt.Println("Função aaaa " + x + "|" + teste)
	}(teste)

	teste = " pedro"

	fmt.Println("hhhhhhhh")
	fmt.Println("########")

	arq, err := LerArquivo("arquivo.csv")
	if err != nil {
		fmt.Println("Erro ao ler o arquivo:", err)
		return
	}
	defer arq.Close()
	defer func() {
		fmt.Println("Função bbbb")
	}()
}
