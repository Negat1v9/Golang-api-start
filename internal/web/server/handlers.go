package server

import (
	"net/http"

	helloservice "github.com/Negat1v9/golang-api-start/internal/services/hello"
	hellohttp "github.com/Negat1v9/golang-api-start/internal/web/hello/http"
	"github.com/Negat1v9/golang-api-start/internal/web/middleware"
)

func (s *Server) initHandlers() {
	// init middleware
	mw := &middleware.MiddleWareManager{}

	// create stack of middleware
	mwStack := middleware.CreateStack(mw.LogRequest)

	// init services
	helloService := helloservice.NewService()
	// init handlers
	helloHandler := hellohttp.NewHandler(helloService)

	baseMux := http.NewServeMux()

	// init routers
	helloRouter := hellohttp.HelloRouter(helloHandler, mw)
	// all requests to the router hello will have the prefix /hello
	baseMux.Handle("/hello/", http.StripPrefix("/hello", helloRouter))

	apiV1Mux := http.NewServeMux()
	// // all requests to the router baseMux will have the prefix /api/v1
	apiV1Mux.Handle("/api/v1/", http.StripPrefix("/api/v1", baseMux))

	// all requests will go through the middleware layer that was introduced in mwStack
	s.server.Handler = mwStack(apiV1Mux)
}
