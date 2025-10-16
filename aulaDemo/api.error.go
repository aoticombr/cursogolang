package main

import (
	"encoding/json"
	"net/http"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, code int, body []byte, err error) {
	if (code >= 200) && (code <= 299) {
		w.WriteHeader(code)
		if body != nil {
			w.Write(body)
		}
	} else {
		w.WriteHeader(code)
		if err != nil {
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		} else {
			w.Write(body)
		}

	}
}
