package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type Pessoa struct {
	Id   string
	Nome string
	Cpf  string
}

func Guid() string {
	id := uuid.New().String()
	return id
}

var Pessoas []*Pessoa

func GetPessoaById(id string) *Pessoa {
	for _, p := range Pessoas {
		if p.Id == id {
			return p
		}
	}
	return nil
}

func GetPessoaByCpf(cpf string) *Pessoa {
	for _, p := range Pessoas {
		if p.Cpf == cpf {
			return p
		}
	}
	return nil
}

func GetPessoaByCpfNotId(cpf, id string) *Pessoa {
	for _, p := range Pessoas {
		if p.Cpf == cpf && p.Id != id {
			return p
		}
	}
	return nil
}

func main() {
	//fmt.Println(Guid())

	r := chi.NewRouter()

	r.Get("/pessoas", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Pessoas)

	})
	r.Post("/pessoa", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var p *Pessoa
		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if GetPessoaByCpf(p.Cpf) != nil {
			http.Error(w, "Pessoa com esse CPF ja existe", http.StatusConflict)
			return
		}

		p.Id = Guid()

		Pessoas = append(Pessoas, p)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(p)

	})
	r.Put("/pessoa/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		id := chi.URLParam(r, "id")

		var p *Pessoa
		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		pessoa := GetPessoaById(id)
		if pessoa == nil {
			http.Error(w, "Pessoa nao encontrada", http.StatusNotFound)
			return
		}
		pessoa2 := GetPessoaByCpfNotId(p.Cpf, id)
		if p.Cpf != pessoa.Cpf && pessoa2 != nil {
			http.Error(w, "Pessoa com esse CPF ja existe", http.StatusConflict)
			return
		}

		pessoa.Nome = p.Nome
		pessoa.Cpf = p.Cpf
	})
	r.Delete("/pessoa/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		id := chi.URLParam(r, "id")

		pessoa := GetPessoaById(id)
		if pessoa == nil {
			http.Error(w, "Pessoa nao encontrada", http.StatusNotFound)
			return
		}

		for i, p := range Pessoas {
			if p.Id == id {
				Pessoas = append(Pessoas[:i], Pessoas[i+1:]...)
				break
			}
		}

		w.WriteHeader(http.StatusNoContent)
	})
	r.Get("/pessoa/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		id := chi.URLParam(r, "id")

		pessoa := GetPessoaById(id)
		if pessoa == nil {
			http.Error(w, "Pessoa nao encontrada", http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(pessoa)
	})

	http.ListenAndServe(":7001", r)
}
