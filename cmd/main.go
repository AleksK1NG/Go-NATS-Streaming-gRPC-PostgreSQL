package main

import (
	"log"
	"net/http"

	"github.com/AleksK1NG/nats-streaming/config"
	"github.com/AleksK1NG/nats-streaming/pkg/jaeger"
	"github.com/AleksK1NG/nats-streaming/pkg/logger"
	"github.com/AleksK1NG/nats-streaming/pkg/redis"
	"github.com/opentracing/opentracing-go"
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

	tracer, closer, err := jaeger.InitJaeger(cfg)
	if err != nil {
		appLogger.Fatal("cannot create tracer", err)
	}
	appLogger.Info("Jaeger connected")

	opentracing.SetGlobalTracer(tracer)
	defer closer.Close()
	appLogger.Info("Opentracing connected")

	client, err := redis.NewRedisClient(cfg)
	if err != nil {
		appLogger.Fatalf("NewRedisClient: %+v", err)
	}

	appLogger.Infof("Redis connected: %+v", client.PoolStats())

	if err := http.ListenAndServe(":5000", nil); err != nil {
		appLogger.Fatalf("ListenAndServe: %+v", err)
	}
}
