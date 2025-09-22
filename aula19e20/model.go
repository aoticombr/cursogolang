package main

type Cliente struct {
	Id        string `db:"id"`
	NomeRazao string `db:"nome_razao"`
}

type Aula19 struct {
	Id   int    `db:"id"`
	Nome string `db:"nome"`
}
