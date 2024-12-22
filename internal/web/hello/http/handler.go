package hellohttp

import (
	"context"
	"encoding/json"
	"net/http"

	helloservice "github.com/Negat1v9/golang-api-start/internal/services/hello"
)

type HelloHandler struct {
	service *helloservice.HelloService
}

func NewHandler(s *helloservice.HelloService) *HelloHandler {
	return &HelloHandler{
		service: s,
	}

}

func (h *HelloHandler) Hello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		res := h.service.Hello(context.Background())

		json.NewEncoder(w).Encode(&res)
	}
}
