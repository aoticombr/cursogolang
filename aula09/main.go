package main

import "fmt"

type Pessoa struct {
	Peso float32
	Cor  string
}

func IncA(value *int) {
	*value++

}

func IncB(value *int) *int {
	*value++
	return value
}

func ModPessoaA(value *Pessoa) {
	value.Cor = "amarelo"
	value.Peso = 300

	value = &Pessoa{Peso: 90, Cor: "Branca"}
}

func ModPessoaB(value *Pessoa) *Pessoa {
	value.Cor = "amarelo"
	value.Peso = 300

	value1 := &Pessoa{Peso: 90, Cor: "Branca"}

	value.Cor = "vinho"
	value.Peso = 10
	return value1
}

func main() {
	var contador *int
	v1 := 0
	contador = &v1
	fmt.Println(*contador)
	contador = IncB(contador)
	fmt.Println(*contador)
	IncA(contador)
	fmt.Println(*contador)
	IncB(contador)
	fmt.Println(*contador)
	IncB(contador)
	fmt.Println(*contador)

	var pessoa *Pessoa
	pessoa = &Pessoa{}
	fmt.Println(*pessoa)

	ModPessoaA(pessoa)

	fmt.Println(*pessoa)

	ModPessoaB(pessoa)

	fmt.Println(*pessoa)

	pessoa = ModPessoaB(pessoa)

	fmt.Println(*pessoa)
}
