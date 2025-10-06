package main

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/aoticombr/golang/lib"
)

type Pessoa struct {
	Nome       string
	Idade      int
	DtaNasc    string
	Logradouro string
	Bairro     string
	Cidade     string
}

func (p *Pessoa) Imprimir() {
	fmt.Printf("%+v\n", p)
}

func ConverteLinhaEmPessoa1(value []string) *Pessoa {
	p := &Pessoa{
		Nome:       value[0],
		Idade:      lib.StrToInt(value[1]),
		DtaNasc:    value[2],
		Logradouro: value[3],
		Bairro:     value[4],
		Cidade:     value[5],
	}
	return p
}
func ConverteLinhaEmPessoa2(value []string) *Pessoa {
	p := &Pessoa{}
	p.Nome = value[0]
	p.Idade = lib.StrToInt(value[1])
	p.DtaNasc = value[2]
	p.Logradouro = value[3]
	p.Bairro = value[4]
	p.Cidade = value[5]
	return p
}

func main() {
	file, err := os.Open("aula4.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = '|'
	reader.ReuseRecord = true //economiza memoria
	reader.TrimLeadingSpace = true

	var Pessoas []*Pessoa

	linha := 0
	for {
		record, err := reader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			panic(err)
		}
		linha++
		if linha == 1 {
			continue
		}
		pessoa := ConverteLinhaEmPessoa2(record)
		pessoa.Imprimir()

		Pessoas = append(Pessoas, pessoa)

	}
	fmt.Print(Pessoas)
}
