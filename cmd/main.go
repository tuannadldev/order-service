package main

import (
	"log"
	"order-service/config"
	"order-service/internal/server"
	"order-service/pkg/logger"
)

func main() {
	cfg, err := config.InitConfig("local")
	if err != nil {
		log.Fatalln("Failed at config", err)
	}
	appLogger := logger.InitAppLogger(&cfg.Logger)
	appLogger.InitLogger()
	appLogger.WithName(cfg.Server.ServiceName)
	appLogger.Fatal(server.InitServer(cfg, appLogger).Run())
}
