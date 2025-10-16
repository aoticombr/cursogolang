package api

import (
	"net/http"

	"github.com/aoticombr/golang/framework/api"
)

func ActionBegin(r *http.Request) {
	println("Metodo:", r.Method, " - URL:", r.URL.String())
}

func ActionEnd(w http.ResponseWriter, r *http.Request, code int, body []byte, err error) {
	println("Status code:", code)
	if body != nil {
		println("Body:", string(body))
	}
	if err != nil {
		println("Erro:", err.Error())
	}
	api.ErrorHandler(w, r, code, body, err)
}
