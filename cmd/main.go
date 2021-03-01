package main

import (
	"log"

	"github.com/AleksK1NG/nats-streaming/config"
	"github.com/AleksK1NG/nats-streaming/pkg/logger"
)

func main() {
	log.Println("Starting microservice")

	cfg, err := config.ParseConfig()
	if err != nil {
		log.Fatal(err)
	}

	appLogger := logger.NewApiLogger(cfg)
	appLogger.InitLogger()
	appLogger.Info("Starting user server")
	appLogger.Infof(
		"AppVersion: %s, LogLevel: %s, DevelopmentMode: %s",
		cfg.AppVersion,
		cfg.Logger.Level,
		cfg.HTTP.Development,
	)
	appLogger.Infof("Success parsed config: %+v", cfg.AppVersion)
}
