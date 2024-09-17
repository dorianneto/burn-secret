package middleware

import (
	"log/slog"
	"net/http"
)

func LogRequests(next http.Handler, logger *slog.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			method = r.Method
			uri    = r.URL.RequestURI()
			ip     = r.RemoteAddr
		)

		logger.Info("received request", "ip", ip, "method", method, "URI", uri)

		next.ServeHTTP(w, r)
	})
}
