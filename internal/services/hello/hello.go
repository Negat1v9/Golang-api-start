package helloservice

import (
	"context"

	modelhello "github.com/Negat1v9/golang-api-start/model/hello"
)

type HelloService struct {
	// you can specify the layer for working with
	// the database and the config for this service
}

func NewService() *HelloService {
	return &HelloService{}
}

func (s *HelloService) Hello(ctx context.Context) *modelhello.HelloRes {
	return &modelhello.HelloRes{
		Data: "hello User! for information go to ",
	}
}
