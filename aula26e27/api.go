package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type CoreApi struct {
	Porta string
}

func (cApi *CoreApi) Start() {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("API no ar!"))

	})

	fmt.Println("API rodando na porta", cApi.Porta)
	http.ListenAndServe(":"+cApi.Porta, r)

}

func NewCoreApi() *CoreApi {
	api := &CoreApi{
		Porta: "9090",
	}
	return api
}
