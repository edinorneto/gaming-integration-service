package main

import (
	"log/slog"
	"net/http"
	"time"
)

// responseWriter registra o status HTTP enviado pela API.
type responseWriter struct {
	http.ResponseWriter
	status int
}

// WriteHeader captura o status antes de enviá-lo para o cliente.
func (w *responseWriter) WriteHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}

// loggingMiddleware registra informações de todas as requisições.
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		start := time.Now()

		writer := &responseWriter{
			ResponseWriter: w,
			status:         http.StatusOK,
		}

		next.ServeHTTP(writer, r)

		slog.Info(
			"Requisição finalizada",
			"method", r.Method,
			"path", r.URL.Path,
			"status", writer.status,
			"duration", time.Since(start),
		)
	})
}