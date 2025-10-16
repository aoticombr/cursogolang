package api

import (
	"net/http"

	fr "github.com/aoticombr/golang/framework"
	api "github.com/aoticombr/golang/framework/api"
)

func RegisterPublicRoutes() {
	RoutePublic.RegisterRoute("/ping", api.GET, func(w http.ResponseWriter, r *http.Request) {
		ActionBegin(r)
		Code, Body, Err := Ping()
		ActionEnd(w, r, Code, Body, Err)
	})
	RoutePublic.RegisterRoute("/login", api.GET, func(w http.ResponseWriter, r *http.Request) {
		ActionBegin(r)
		Code, Body, Err := Login(w, r, fr.CoreApi)
		ActionEnd(w, r, Code, Body, Err)
	})
}
