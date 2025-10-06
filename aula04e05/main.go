package main

import (
	"fmt"
	"math/rand"
	"time"

	cp "github.com/aoticombr/golang/component"
	lib "github.com/aoticombr/golang/lib"
)

func SortearIdade(ate int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(ate)
}

func Cabecalho() string {
	return "Nome|Idade|Dta Nasci|Logradouro|Bairro|Cidade"
}
func LinhaModelo1(i int) string {
	Idade := SortearIdade(50)
	Data := time.Now()
	Data = Data.AddDate(0, 0, -Idade)

	//forma 1
	linha := "Nome " + lib.IntToStr(i)
	linha += "|" + lib.IntToStr(Idade)
	linha += "|" + Data.Format("02/01/2006")
	linha += "|Rua " + lib.IntToStr(i)
	linha += "|Bairro " + lib.IntToStr(i)
	linha += "|Cidade " + lib.IntToStr(i)
	return linha
}

func LinhaModelo2(i int) string {
	Idade := SortearIdade(50)
	Data := time.Now()
	Data = Data.AddDate(0, 0, -Idade)

	//forma 2
	linha := "Nome " + lib.IntToStr(i) + "|" + lib.IntToStr(Idade) + "|" + Data.Format("02/01/2006") + "|Rua " + lib.IntToStr(i) + "|Bairro " + lib.IntToStr(i) + "|Cidade " + lib.IntToStr(i)

	return linha
}

func LinhaModelo3(i int) string {
	Idade := SortearIdade(50)
	Data := time.Now()
	Data = Data.AddDate(0, 0, -Idade)

	//forma 3
	return fmt.Sprintf("Nome %d|%d|%s|Rua %d|Bairro %d|Cidade %d", i, Idade, Data.Format("02/01/2006"), i, i, i)
}

func CriarArquivoGb(value int) {
	var arquivo cp.Strings
	arquivo.Delimiter = "\n"

	arquivo.Add(Cabecalho())

	for i := 0; i < 10*value; i++ {
		arquivo.Add(LinhaModelo3(i))
	}

	err := lib.ByteToSaveFile("../aula7/aula4.txt", arquivo.Byte())
	if err != nil {
		fmt.Println("erro ao gravar o arquivo", err)
	}
}

func main() {
	CriarArquivoGb(1)
}
