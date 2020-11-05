package common

import (
	"net/http"

	"github.com/rs/zerolog/log"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Our middleware logic goes here...
		log.Info().Msg("Request")
		next.ServeHTTP(w, r)
	})
}
