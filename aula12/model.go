package main

import "time"

type Item struct {
	Descricao     string  `xml:"descricao"`
	Quantidade    int     `xml:"quantidade"`
	ValorUnitario float64 `xml:"valor_unitario"`
	Observacao    *string `xml:"observacao,omitempty"`
}

type Nota struct {
	XmlAtrr    string    `xml:"tipo,attr"`
	DtaCompra  time.Time `xml:"dta_compra"`
	TotalNota  float64   `xml:"total_nota"`
	Numero     int       `xml:"numero"`
	Cliente    *string   `xml:"cliente"`
	Logradouro string    `xml:"logradouro,omitempty"`
	Itens      []Item    `xml:"itens"`
}

type Notas struct {
	Notas []Nota `xml:"notas"`
}
