package api

import (
	"net/http"

	"github.com/aoticombr/golang/framework/api"
)

func RegisterPrivateRoutes() {
	RoutePrivate.RegisterRoute("/pingpriv", api.GET, func(w http.ResponseWriter, r *http.Request) {
		ActionBegin(r)
		Code, Body, Err := Ping()
		ActionEnd(w, r, Code, Body, Err)
	})
}
