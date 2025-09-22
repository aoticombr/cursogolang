package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var Pessoas []*Pessoa

const (
	apiKey   = "14jk1l4j1lj41k4343"
	username = "admin"
	password = "12345"

	jwtSecret = "minha-chave"

	clientID     = "meu-client-id"
	clientSecret = "meu-client-secret"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type OAuthLoginRequest struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	GrantType    string `json:"grant_type"`
}

type OAuthLoginResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}

type Pessoa struct {
	Id   string
	Nome string
	Cpf  string
}

func Guid() string {
	id := uuid.New().String()
	return id
}

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

// autenticaçao api-key
func apiKeyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key := r.Header.Get("X-API-Key")
		if key != apiKey {
			http.Error(w, "API KEY INVÁLIDA", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// autenticacao JWT
func jwtMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "token ausente", http.StatusUnauthorized)
			return
		}
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Metodo de assinatura inválido")
			}
			return []byte(jwtSecret), nil
		})
		if err != nil || !token.Valid {
			http.Error(w, "Token inválido", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)

	})
}

func oauthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if !strings.HasPrefix(auth, "Bearer ") {
			http.Error(w, "Token ausente", http.StatusUnauthorized)
			return
		}
		tokenString := auth[7:]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Método de assinatura inválido")
			}
			return []byte(jwtSecret), nil
		})
		if err != nil || !token.Valid {
			http.Error(w, "Token inválido", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

//login jwt

func loginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "JSON invalidas", http.StatusUnauthorized)
		return
	}
	claims := jwt.MapClaims{
		"username": req.Username,
		"exp":      time.Now().Add(time.Hour * 1).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		http.Error(w, "Erro ao gerar token", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(LoginResponse{Token: tokenString})

}

func oauthLoginHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Erro ao processar form-data", http.StatusBadRequest)
		return
	}
	clientIDForm := r.FormValue("client_id")
	clientSecretForm := r.FormValue("client_secret")
	grantTypeForm := r.FormValue("grant_type")

	if grantTypeForm != "client_credentials" {
		http.Error(w, "grant_type inválido ou ausente", http.StatusBadRequest)
		return
	}
	if clientIDForm != clientID || clientSecretForm != clientSecret {
		http.Error(w, "Credenciais inválidas", http.StatusUnauthorized)
		return
	}
	claims := jwt.MapClaims{
		"client_id": clientIDForm,
		"exp":       time.Now().Add(time.Hour * 1).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		http.Error(w, "Erro ao gerar token", http.StatusInternalServerError)
		return
	}
	resp := OAuthLoginResponse{
		AccessToken: tokenString,
		TokenType:   "Bearer",
	}
	json.NewEncoder(w).Encode(resp)
}

func main() {
	//fmt.Println(Guid())

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	//r.Use(apiKeyMiddleware)
	//r.Post("/login", loginHandler)
	r.Post("/login", oauthLoginHandler)
	r.With(oauthMiddleware).Get("/pessoas", func(w http.ResponseWriter, r *http.Request) {
		//panic("teste")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Pessoas)

	})
	r.With(oauthMiddleware).Post("/pessoa", func(w http.ResponseWriter, r *http.Request) {
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
	r.With(oauthMiddleware).Put("/pessoa/{id}", func(w http.ResponseWriter, r *http.Request) {
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
	r.With(oauthMiddleware).Delete("/pessoa/{id}", func(w http.ResponseWriter, r *http.Request) {
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
	r.With(oauthMiddleware).Get("/pessoa/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		id := chi.URLParam(r, "id")

		pessoa := GetPessoaById(id)
		if pessoa == nil {
			http.Error(w, "Pessoa nao encontrada", http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(pessoa)
	})
	fmt.Println("http://localhost:7001")
	http.ListenAndServe(":7001", r)
}
