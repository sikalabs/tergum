package server

import (
	"github.com/rs/zerolog/log"
)

func logServerStarted(addr string) {
	log.Info().
		Str("log-id", "server-started").
		Str("addr", addr).
		Msg("Server started on addr: " + addr + ".")
}

func logApiCall(method string, path string) {
	log.Info().
		Str("log-id", "server-api-call").
		Str("method", method).
		Str("path", path).
		Msg(method + " " + path)
}
