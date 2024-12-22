package hellohttp

import (
	"net/http"

	"github.com/Negat1v9/golang-api-start/internal/web/middleware"
)

func HelloRouter(h *HelloHandler, mw *middleware.MiddleWareManager) http.Handler {
	helloMux := http.NewServeMux()

	helloMux.HandleFunc("GET /hello", h.Hello())

	router := http.NewServeMux()
	// to gain access to the helloMux router, the request is checked in
	// the middleware to authorize the user
	router.Handle("/", mw.AuthUser(helloMux))

	return router
}
