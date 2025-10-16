package main

import (
	"fmt"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request, CoreApi *CoreApi) (int, []byte, error) {
	w.Header().Set("Content-Type", "application/json")
	ok, jwt := Config.GetJwt("admin")
	if !ok {
		return 500, nil, fmt.Errorf("Configuração JWT 'admin' não encontrada")
	}

	// Ds := db.NewDataSet()
	// defer Ds.Close()
	// Ds.Sql.Clear();
	// Ds.Sql.Add("SELECT 1 AS ID, 'Usuário Logado' AS MSG")
	return 200, nil, nil
}
