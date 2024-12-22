package main

import (
	"log"

	"github.com/Negat1v9/golang-api-start/internal/config"
	"github.com/Negat1v9/golang-api-start/internal/web/server"
	"github.com/Negat1v9/golang-api-start/pkg/logger"
)

func main() {
	// you can initialize the client for the database, create a config, and so on

	// set default logger for app
	logger.InitLogger()

	// at the moment the config is a stub
	cfg := config.NewConfig()

	// init server
	server := server.NewServer(cfg)

	if err := server.Run(); err != nil {
		log.Println("server is stopped")
	}
}
