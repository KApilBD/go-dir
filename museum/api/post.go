package api

import (
	"encoding/json"
	"net/http"

	"fm.com/go/museum/data"
)

func Post(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var exhibition data.Exhibition
		err := json.NewDecoder(r.Body).Decode(&exhibition)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		data.Add(exhibition)
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("Ok"))
	} else {
		http.Error(w, "Unsuported Method", http.StatusMethodNotAllowed)
	}
}
