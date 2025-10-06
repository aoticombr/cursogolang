package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type CoreApi struct {
	Porta string
}

func (cApi *CoreApi) Start() error {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("API no ar!"))
	})
	fmt.Println("API iniciada na porta " + cApi.Porta)
	// Inicie o servidor
	http.ListenAndServe(":"+cApi.Porta, r)

	return nil
}

func NewCoreApi() *CoreApi {
	api := &CoreApi{
		Porta: "9090",
	}

	return api
}
