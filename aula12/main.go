package main

import (
	"encoding/xml"
	"fmt"
	"time"
)

func main() {
	var notas Notas
	var nota Nota
	nota.XmlAtrr = "YYYY-MM-DD"
	nota.DtaCompra = time.Now()
	nota.TotalNota = 150.75
	nota.Numero = 12345
	nome := "Cliente XYZ"
	nota.Cliente = &nome

	nota.Itens = []Item{
		{Descricao: "Item A", Quantidade: 2, ValorUnitario: 25.00},
		{Descricao: "Item B", Quantidade: 1, ValorUnitario: 100.75},
	}

	var item Item
	item.Descricao = "Item C"
	item.Quantidade = 3
	item.ValorUnitario = 50.00
	obs := "teste"
	item.Observacao = &obs

	nota.Itens = append(nota.Itens, item)

	notas.Notas = append(notas.Notas, nota)
	notas.Notas = append(notas.Notas, nota)

	fmt.Println("##################################################")
	fmt.Println("Modelo 1")
	BodyXml, err := xml.MarshalIndent(nota, "", "   ")
	if err != nil {
		panic(err)
	}
	println(string(BodyXml))

}
