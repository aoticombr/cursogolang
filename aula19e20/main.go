package main

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
)

const PgUser = "postgres"
const PgPass = "master"
const PgServer = "localhost"
const PgPort = 5432
const PgDb = "postgres"
const Url = "postgres://%s:%s@%s:%d/%s"

func main() {
	DSN := fmt.Sprintf(Url, PgUser, PgPass, PgServer, PgPort, PgDb)

	conn, err := sql.Open("pgx", DSN)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	var result int
	err = conn.QueryRow("SELECT 1000").Scan(&result)
	if err != nil {
		panic(err)
	}
	fmt.Println("resultado da query:", result)

	fmt.Println("======================1=============================")
	var clientes []Cliente
	rows, err := conn.Query("SELECT id, nome_razao FROM public.tbcli_0001 limit 10")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var c Cliente
		err := rows.Scan(&c.Id, &c.NomeRazao)
		if err != nil {
			panic(err)
		}
		clientes = append(clientes, c)
	}

	for _, cliente := range clientes {
		fmt.Printf("ID: %s, Nome/Razao: %s\n", cliente.Id, cliente.NomeRazao)
	}
	fmt.Println("======================2=============================")
	var Aula19s []Aula19
	rows, err = conn.Query("SELECT id, nome FROM public.aula19")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var c Aula19
		err := rows.Scan(&c.Id, &c.Nome)
		if err != nil {
			panic(err)
		}
		Aula19s = append(Aula19s, c)
	}
	fmt.Println("======================2.1=============================")
	for _, aula := range Aula19s {
		fmt.Printf("ID: %d, Nome: %s\n", aula.Id, aula.Nome)
	}
	fmt.Println("======================2.2=============================")
	for _, aula := range Aula19s {
		fmt.Printf("ID: %d, Nome: %s\n", aula.Id, aula.Nome)
	}
	fmt.Println("======================2.3=============================")
	for _, aula := range Aula19s {
		fmt.Printf("ID: %d, Nome: %s\n", aula.Id, aula.Nome)
	}
	fmt.Println("======================3=============================")
	var Id int
	var Nome string

	rows, err = conn.Query("SELECT id, nome FROM public.aula19")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {

		err := rows.Scan(&Id, &Nome)
		if err != nil {
			panic(err)
		}
		fmt.Printf("ID: %d, Nome: %s\n", Id, Nome)
	}
	var count int
	err = conn.QueryRow("SELECT max(id) as id FROM public.aula19").Scan(&count)
	if err != nil {
		panic(err)
	}

	for i := count + 1; i <= count+3; i++ {
		_, err := conn.Exec("insert into public.aula19(id,nome) values($1,$2)", i, fmt.Sprintf("Nome %d", i))
		if err != nil {
			panic(err)
		}
	}
	_, err = conn.Exec("update public.aula19 set nome=$1 where id=$2", "teste 1", 6)
	if err != nil {
		panic(err)
	}
	_, err = conn.Exec("update public.aula19 set nome=$1 where id=$2", "teste 2", 4)
	if err != nil {
		panic(err)
	}
	_, err = conn.Exec("delete from public.aula19 where id=$1", 7)
	if err != nil {
		panic(err)
	}
}
