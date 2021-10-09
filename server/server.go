package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sikalabs/tergum/version"
)

func response(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(data)
}

func Server(addr string) {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		response(w, map[string]interface{}{
			"meta": map[string]interface{}{
				"name":    "tergum",
				"version": version.Version,
			},
		})
	})
	fmt.Println("Server started.")
	http.ListenAndServe(addr, router)
}
