package middleware

import (
	"log"
	"net/http"
	"time"
)

func (m *MiddleWareManager) LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("Request: From: %s, Method: %s, Time: %s", r.RemoteAddr, r.Method, time.Since(start).String())
	})
}
