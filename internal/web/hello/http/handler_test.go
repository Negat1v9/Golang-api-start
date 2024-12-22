package hellohttp_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	helloservice "github.com/Negat1v9/golang-api-start/internal/services/hello"
	hellohttp "github.com/Negat1v9/golang-api-start/internal/web/hello/http"
	"github.com/Negat1v9/golang-api-start/internal/web/middleware"
)

// tests for all hello handlers, just an example, in the project you need
// to make mocks for services to take into account all possible error options

func TestHello(t *testing.T) {
	helloService := helloservice.NewService()
	helloHandler := hellohttp.NewHandler(helloService)

	mw := &middleware.MiddleWareManager{}
	handler := hellohttp.HelloRouter(helloHandler, mw)

	req, err := http.NewRequest("GET", "/hello", nil)
	if err != nil {
		t.Errorf("create request %v", err)
		return
	}

	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	if rr.Code != 200 {
		t.Errorf("expected %d received %d", 200, rr.Code)
	}
}
