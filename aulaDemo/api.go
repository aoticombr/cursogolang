package main

import (
	"fmt"
	"net/http"

	"github.com/aoticombr/golang/config"
	"github.com/aoticombr/golang/dbconndataset"
	"github.com/go-chi/chi/v5"
)

type CoreApi struct {
	Porta string
	Conf  *config.Api
	Conn  *dbconndataset.ConnDataSet
}

func (cApi *CoreApi) Start() {
	r := chi.NewRouter()

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("API no ar!"))

	})

	r.Get("/login", func(w http.ResponseWriter, r *http.Request) {
		ActionBegin(w, r)
		Status, Body, Err := Login(w, r, cApi)
		ActionEnd(w, r, Status, Body, Err)
	})

	fmt.Println("API rodando na porta", cApi.Porta)
	http.ListenAndServe(":"+cApi.Porta, r)

}

func NewCoreApi() *CoreApi {
	var err error
	var ok bool
	api := &CoreApi{}
	ok, api.Conf = Config.GetApi(NomeApp)
	if !ok {
		panic("Configuração da API não encontrada")
	}
	api.Porta = api.Conf.GetPortStr()
	if !ok {
		panic("Configuração da porta da API não encontrada")
	}
	ok, DBName := Config.GetDB(api.Conf.Dbs[0])
	if !ok {
		panic("Configuração do banco de dados não encontrada")
	}
	api.Conn, err = dbconndataset.NewConn(*DBName)
	if err != nil {
		panic("erro ao conectar ao banco dados: " + err.Error())
	}

	return api
}
