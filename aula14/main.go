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

	r.Post("/pessoa", func(w http.ResponseWriter, r *http.Request) {
		var p Pessoa
		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		p.Idade += 1 // Incrementa a idade em 1 ano como exemplo de processamento
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(p)
	})

	r.Put("/pessoa/{id}", func(w http.ResponseWriter, r *http.Request) {
		IdStr := chi.URLParam(r, "id")
		var p Pessoa
		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		demitido := r.URL.Query().Get("demitido") // Exemplo de como acessar query params, se necessário
		if demitido == "" {
			http.Error(w, "Query param 'demitido' is required", http.StatusBadRequest)
			return
		}

		HeaderApiKey := r.Header.Get("x-api-key") // Exemplo de como acessar headers, se necessário
		if HeaderApiKey == "" {
			http.Error(w, "Header 'x-api-key' is required", http.StatusBadRequest)
			return
		}
		if HeaderApiKey != "cursogolang" {
			http.Error(w, "Invalid 'x-api-key' header", http.StatusUnauthorized)
			return
		}

		p.Nome = p.Nome + " " + IdStr + " " + demitido // Atualiza o nome como exemplo de processamento
		p.Idade += 1                                   // Incrementa a idade em 1 ano como exemplo de processamento
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(p)
	})

	http.ListenAndServe(":7001", r)
}
