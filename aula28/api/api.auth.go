package api

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/aoticombr/golang/framework/api"
	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt"
)

func Login(w http.ResponseWriter, r *http.Request, core *api.CoreApi) (int, []byte, error) {
	w.Header().Set("Content-Type", "application/json")
	db := chi.URLParam(r, "db")
	ts := core.Processo.Tenants.FindByName(db)
	if ts == nil {
		return http.StatusNotFound, nil, nil
	}
	/*Basic sadkaslkdjalksjdalksjdakjsdljalsdjkal
	  [0]Basic [1]dfsfsdfsdf
	  Bearer sdkfjsdkfjsdlfksdlkfjsld*/
	authHeader := r.Header.Get("Authorization") //Basic.  usuario:senha. === base64
	if authHeader == "" {
		return http.StatusNotFound, nil, nil
	}
	basic := strings.Split(authHeader, " ")
	if len(basic) != 2 || basic[0] != "Basic" {
		return http.StatusUnauthorized, nil, nil
	}
	credential, err := base64.StdEncoding.DecodeString(basic[1])
	if err != nil {
		return http.StatusUnauthorized, nil, nil
	}
	userpass := strings.Split(string(credential), ":")
	user := userpass[0]
	pass := userpass[1]
	//estudar SQL INJECTOR

	hashPass := md5.Sum([]byte(pass))
	passMd5 := hex.EncodeToString(hashPass[:])

	Ds := ts.Connection.NewDataSet()
	defer Ds.Free()
	Ds.Sql.Clear()
	Ds.Sql.Add(`select id from usuarios where login = :login and senha = :senha`)
	Ds.SetInputParam("login", user)
	Ds.SetInputParam("senha", passMd5)
	err = Ds.Open()
	if err != nil {
		return http.StatusInternalServerError, nil, nil
	}
	if Ds.Count() > 1 {
		return http.StatusUnauthorized, nil, nil
	}
	if Ds.Count() == 0 {
		return http.StatusUnauthorized, nil, nil
	}
	if Ds.Count() == 1 {
		exp := time.Now().Add(time.Hour * 1).Unix()
		claim := jwt.MapClaims{
			"id":  Ds.FieldByName("id").AsString(),
			"exp": exp,
		}
		_, TokenString, err := TokenAuth.Encode(claim)
		if err != nil {
			return http.StatusInternalServerError, nil, nil
		}
		response := map[string]string{"token": TokenString, "exp": fmt.Sprintf("%d", exp)}
		json.NewEncoder(w).Encode(response)
		return http.StatusOK, nil, nil

	} else {
		return http.StatusUnauthorized, nil, nil
	}

}
