package server

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/sikalabs/tergum/backup/source"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/sikalabs/tergum/version"
)

func response(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(data)
}

func Server(addr string) {
	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339,
	})

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		response(w, map[string]interface{}{
			"meta": map[string]interface{}{
				"name":    "tergum",
				"version": version.Version,
			},
		})
		logApiCall("GET", "/")
	})
	router.HandleFunc("/api/v1/backup", func(w http.ResponseWriter, r *http.Request) {
		var source source.Source

		body, _ := ioutil.ReadAll(r.Body)
		_ = r.Body.Close()
		_ = json.Unmarshal(body, &source)

		data, _, _ := source.Backup()

		w.Header().Set("Content-Type", "application/octet-stream")
		io.Copy(w, data)

		logApiCall("POST", "/api/v1/backup")
	})
	logServerStarted(addr)
	http.ListenAndServe(addr, router)
}
