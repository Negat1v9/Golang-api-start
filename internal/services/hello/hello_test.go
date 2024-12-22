package helloservice_test

import (
	"context"
	"testing"

	helloservice "github.com/Negat1v9/golang-api-start/internal/services/hello"
	modelhello "github.com/Negat1v9/golang-api-start/model/hello"
)

// tests for all HelloService methods, just an example, in the project you need
// to make mocks for services to take into account all possible error options

func TestHello(t *testing.T) {
	service := helloservice.NewService()
	res := service.Hello(context.Background())

	testResult := modelhello.HelloRes{
		Data: "hello User! for information go to https://github.com/Negat1v9/golang-api-start",
	}
	if testResult.Data != res.Data {
		t.Errorf("TestHello must %s, get %s", testResult.Data, res.Data)
	}
}
