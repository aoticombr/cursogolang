package api

import (
	"net/http"

	config "github.com/aoticombr/golang/config"
	api "github.com/aoticombr/golang/framework/api"
	"github.com/go-chi/jwtauth"
)

var (
	TokenAuth *jwtauth.JWTAuth

	RoutePublic  *api.RouteGroup //rota publica
	RoutePrivate *api.RouteGroup //rota privada (autenticada)
)

func DynamicVerifier(authFunc func() *jwtauth.JWTAuth) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			auth := authFunc()
			verifier := jwtauth.Verifier(auth)
			authenticator := jwtauth.Authenticator

			verifier(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				authenticator(next).ServeHTTP(w, r)
			})).ServeHTTP(w, r)
		})
	}
}

func init() {
	ok, value := config.NewConfig().GetJwt("ADMIN")
	if ok {
		TokenAuth = jwtauth.New("HS256", []byte(value.SecretKey), nil)
	} else {
		panic("Chave [JWT]Admin náo encontrada")
	}

	RoutePublic = api.RegisterRouterGroup("public")

	RoutePrivate = api.RegisterRouterGroup("private")
	RoutePrivate.Use(DynamicVerifier(func() *jwtauth.JWTAuth { return TokenAuth }))

	RegisterPublicRoutes()
	RegisterPrivateRoutes()
}
