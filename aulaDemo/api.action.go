package main

import (
	"net/http"
)

func ActionBegin(w http.ResponseWriter, r *http.Request) {
	println("Método:", r.Method, "URL:", r.URL.String())
}

func ActionEnd(w http.ResponseWriter, r *http.Request, code int, body []byte, err error) {
	println("Resultado Final:")
	println("Status Code:", code)
	if body != nil {
		println("Body:", string(body))
	}
	if err != nil {
		println("Erro:", err.Error())
	}
	ErrorHandler(w, r, code, body, err)
}
