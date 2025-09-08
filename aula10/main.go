package main

import (
	"encoding/json"
	"fmt"
	"time"
)

func main() {
	var notas Notas
	var notas2 []Nota
	var nota Nota
	nota.DtaCompra = time.Now()
	nota.TotalNota = 150.75
	nota.Numero = 12345

	nota.Itens = []Item{
		{Descricao: "Item A", Quantidade: 2, ValorUnitario: 25.00},
		{Descricao: "Item B", Quantidade: 1, ValorUnitario: 100.75},
	}

	var item Item
	item.Descricao = "Item C"
	item.Quantidade = 3
	item.ValorUnitario = 50.00

	nota.Itens = append(nota.Itens, item)
	fmt.Println("##################################################")
	fmt.Println("Modelo 1")
	BodyJson, err := json.Marshal(nota)
	if err != nil {
		panic(err)
	}
	notas.Notas = append(notas.Notas, nota)
	notas.Notas = append(notas.Notas, nota)

	notas2 = append(notas2, nota)
	notas2 = append(notas2, nota)

	fmt.Println("JSON len:", len(BodyJson))
	println(string(BodyJson))
	fmt.Println("##################################################")
	fmt.Println("Modelo 2")
	BodyJson, err = json.MarshalIndent(nota, "", "   ")
	if err != nil {
		panic(err)
	}
	fmt.Println("JSON len:", len(BodyJson))
	println(string(BodyJson))

	fmt.Println("##################################################")
	fmt.Println("Modelo 3")
	BodyJson, err = json.MarshalIndent(notas, "", "   ")
	if err != nil {
		panic(err)
	}
	fmt.Println("JSON len:", len(BodyJson))
	println(string(BodyJson))
	fmt.Println("##################################################")
	fmt.Println("Modelo 4")
	BodyJson, err = json.Marshal(notas2)
	if err != nil {
		panic(err)
	}
	fmt.Println("JSON len:", len(BodyJson))
	println(string(BodyJson))
	fmt.Println("##################################################")
	var JsonStr = `{
		"dta_compra":"2025-09-08T05:59:40.197144-04:00",
		"total_nota":150.75,
		"numero":12345,
		"itens":[
			{
				"descricao":"Item A",
				"quantidade":2,
				"valor_unitario":25
			},
			{
				"descricao":"Item B",
				"quantidade":1,
				"valor_unitario":100.75
			},
			{
				"descricao":"Item C",
				"quantidade":3,
				"valor_unitario":50
			}
		]
		}`

	fmt.Println("JSON String:", JsonStr)
}
