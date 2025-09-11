package main

import "time"

type Item struct {
	Descricao     string  `json:"descricao"`
	Quantidade    int     `json:"quantidade"`
	ValorUnitario float64 `json:"valor_unitario"`
	Observacao    *string `json:"observacao,omitempty"`
}

type Nota struct {
	DtaCompra  time.Time `json:"dta_compra"`
	TotalNota  float64   `json:"total_nota"`
	Numero     int       `json:"numero"`
	Cliente    *string   `json:"cliente"`
	Logradouro string    `json:"logradouro,omitempty"`
	Itens      []Item    `json:"itens"`
}

type Notas struct {
	Notas []Nota `json:"notas"`
}
