package main

import "time"

type Item struct {
	Descricao     string  `json:"descricao"`
	Quantidade    int     `json:"quantidade"`
	ValorUnitario float64 `json:"valor_unitario"`
}

type Nota struct {
	DtaCompra time.Time `json:"dta_compra"`
	TotalNota float64   `json:"total_nota"`
	Numero    int       `json:"numero"`
	Itens     []Item    `json:"itens"`
}

type Notas struct {
	Notas []Nota `json:"notas"`
}
