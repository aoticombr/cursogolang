package main

import (
	"encoding/json"
	"encoding/xml"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Pessoa struct {
	Nome  string `json:"nome" xml:"nome" db:"nome"`
	Idade int    `json:"idade" xml:"idade" db:"idade"`
}

func main() {

	r := chi.NewRouter()

	r.Get("/json", func(w http.ResponseWriter, r *http.Request) {
		p := Pessoa{Nome: "João", Idade: 30}
		jsonData, _ := json.Marshal(p)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)

	})

	r.Get("/xml", func(w http.ResponseWriter, r *http.Request) {
		p := Pessoa{Nome: "João", Idade: 30}
		xmlData, _ := xml.Marshal(p)
		w.Header().Set("Content-Type", "application/xml")
		w.Write(xmlData)

	})

	http.ListenAndServe(":7001", r)
}
