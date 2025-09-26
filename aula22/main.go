package main

import (
	"fmt"

	_ "github.com/sijms/go-ora/v2"

	dataset "github.com/fitlcarlos/godataset"
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

	conn, err := dataset.NewConnection(dataset.DialectType(dataset.ORACLE), DSN)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	fmt.Println("======================1=============================")
	Ds := dataset.NewDataSet(conn)
	defer Ds.Free()
	Ds.AddSql(`SELECT cdcliente, substr(nome,1,5) as nome FROM cliente where rownum <= 10`)
	err = Ds.Open()
	if err != nil {
		panic(err)
	}
	Ds.First()
	for !Ds.Eof() {
		//fmt.Println(Ds.FieldByName("cdcliente").AsString(), Ds.FieldByName("nome").AsString())

		fmt.Printf("ID: %d, Nome/Razao: %s\n", Ds.FieldByName("cdcliente").AsInt64(), Ds.FieldByName("nome").AsString())
		Ds.Next()
	}

	DsExec := dataset.NewDataSet(conn)
	defer DsExec.Free()
	DsExec.AddSql("update cliente set nome = :nome where cdcliente = :cdcliente")
	DsExec.SetInputParam("nome", "3 teste 2")
	DsExec.SetInputParam("cdcliente", 6495)
	_, err = DsExec.Exec()
	if err != nil {
		panic(err)
	}
	fmt.Println("======================2=============================")
	Ds = dataset.NewDataSet(conn)
	defer Ds.Free()
	Ds.AddSql(`SELECT cdcliente, substr(nome,1,5) as nome, :teste as campo1 FROM cliente where rownum <= 15`)
	Ds.SetInputParam("teste", "oi")
	err = Ds.Open()
	if err != nil {
		panic(err)
	}
	Ds.First()
	for !Ds.Eof() {
		//fmt.Println(Ds.FieldByName("cdcliente").AsString(), Ds.FieldByName("nome").AsString())

		fmt.Printf("ID: %d, Nome/Razao: %s\n",
			Ds.FieldByName("cdcliente").AsInt64(),
			Ds.FieldByName("nome").AsString())
		Ds.Next()
	}
	fmt.Println("======================3=============================")
	fmt.Println(Ds.Count())
	Ds.First()
	fmt.Println(Ds.FieldByName("cdcliente").AsInt64())
	Ds.Next()
	fmt.Println(Ds.FieldByName("cdcliente").AsInt64())
	Ds.Next()
	Ds.Next()
	fmt.Println(Ds.FieldByName("cdcliente").AsByte())
	fmt.Println(Ds.FieldByName("campo1").AsString())

	fmt.Println(Ds.SqlParam())

	var clientes []Cliente
	jsonByte, err := Ds.ToStructJson(&clientes)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(jsonByte))
}
