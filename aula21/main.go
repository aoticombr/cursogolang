package main

import (
	"database/sql"
	"fmt"

	_ "github.com/sijms/go-ora/v2"
)

const OraUser = "NEXUS"
const OraPass = "b0b5371e1a0c24a3f08dedbbffb9df"
const OraServer = "192.168.100.5"
const OraPort = 1521
const OraDb = "XE"

const Url = "oracle://%s:%s@%s:%d/%s"

//21C, ORACLE 11G

type Cliente struct {
	Id   int64
	Nome string
}

func main() {
	DSN := fmt.Sprintf(Url, OraUser, OraPass, OraServer, OraPort, OraDb)

	conn, err := sql.Open("oracle", DSN)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	fmt.Println("======================1=============================")
	var clientes []Cliente
	rows, err := conn.Query("SELECT cdcliente, substr(nome,1,5) as nome FROM cliente where rownum <= 10")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var c Cliente
		err := rows.Scan(&c.Id, &c.Nome)
		if err != nil {
			panic(err)
		}
		clientes = append(clientes, c)
	}

	for _, cliente := range clientes {
		fmt.Printf("ID: %d, Nome/Razao: %s\n", cliente.Id, cliente.Nome)
	}

	_, err = conn.Exec("update cliente set nome=:1 where cdcliente=:2", "1 teste 1", 6495)
	if err != nil {
		panic(err)
	}
	_, err = conn.Exec("update cliente set nome=:1 where cdcliente=:2", "2 teste 2", 6479)
	if err != nil {
		panic(err)
	}
	/*_, err = conn.Exec("delete from cliente where cdcliente=:1", 6614)
	if err != nil {
		panic(err)
	}*/

}
