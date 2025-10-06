package main

import "fmt"

type Pessoa struct {
	Peso float32
	Cor  string
}

func main() {

	var (
		//string, float32, int
		v1 string
		v2 *string
		//----------
		v3 Pessoa
		v4 *Pessoa
	)
	v1 = "pedro"

	nome := "zezinho"
	v2 = &nome

	v3.Cor = "vermelho"
	v3.Peso = 1.68

	// v4 = &Pessoa{
	// 	Peso: 2.10,
	// 	Cor:  "Branco",
	// }

	fmt.Println("=======begin========")
	fmt.Println(fmt.Sprintf("v1 = %v", v1))
	fmt.Println("========1========")
	fmt.Println(fmt.Sprintf("v2 = %v", *v2))
	fmt.Println("========2========")
	fmt.Println(fmt.Sprintf("v3 = %v", v3))
	fmt.Println("========4========")
	fmt.Println(fmt.Sprintf("v4.Peso = %v", v3.Peso))
	fmt.Println("========5========")
	fmt.Println(fmt.Sprintf("v4.Cor = %v", v3.Cor))
	fmt.Println("========3========")
	fmt.Println(fmt.Sprintf("v4 = %v", v4))

	if v4 != nil {
		fmt.Println("========4========")
		fmt.Println(fmt.Sprintf("v4.Peso = %v", v4.Peso))
		fmt.Println("========5========")
		fmt.Println(fmt.Sprintf("v4.Cor = %v", v4.Cor))
	}

	fmt.Println("========end========")
}
