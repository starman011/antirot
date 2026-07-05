// Package controller is the HTTP layer: routing, validation, serialization only.
package controller

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/starman011/antirot/backend/internal/service"
)

func NewRouter(sessions *service.SessionService) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/v1/health", handleHealth)

	sc := &sessionController{sessions: sessions}
	mux.HandleFunc("GET /api/v1/session/piece", sc.handlePiece)

	return recoverer(securityHeaders(requestLogger(mux)))
}

func requestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		slog.Info("http", "method", r.Method, "path", r.URL.Path, "dur", time.Since(start))
	})
}

func securityHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h := w.Header()
		h.Set("X-Content-Type-Options", "nosniff")
		h.Set("X-Frame-Options", "DENY")
		h.Set("Referrer-Policy", "no-referrer")
		next.ServeHTTP(w, r)
	})
}

// recoverer logs panics loudly and returns 500, so one request cannot kill the process.
func recoverer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				slog.Error("panic", "recovered", rec, "path", r.URL.Path)
				writeError(w, http.StatusInternalServerError, "internal", "something went wrong")
			}
		}()
		next.ServeHTTP(w, r)
	})
}
