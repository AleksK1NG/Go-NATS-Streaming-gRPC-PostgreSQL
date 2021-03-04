package main

import (
	"log"

	"github.com/AleksK1NG/nats-streaming/config"
	"github.com/AleksK1NG/nats-streaming/internal/server"
	"github.com/AleksK1NG/nats-streaming/pkg/jaeger"
	"github.com/AleksK1NG/nats-streaming/pkg/logger"
	"github.com/AleksK1NG/nats-streaming/pkg/nats"
	"github.com/AleksK1NG/nats-streaming/pkg/postgresql"
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

	redisClient, err := redis.NewRedisClient(cfg)
	if err != nil {
		appLogger.Fatalf("NewRedisClient: %+v", err)
	}

	appLogger.Infof("Redis connected: %+v", redisClient.PoolStats())

	natsConn, err := nats.NewNatsConnect(cfg, appLogger)
	if err != nil {
		appLogger.Fatalf("NewRedisClient: %+v", err)
	}
	appLogger.Infof(
		"Nats Connected: Status: %+v IsConnected: %v ConnectedUrl: %v ConnectedServerId: %v",
		natsConn.NatsConn().Status(),
		natsConn.NatsConn().IsConnected(),
		natsConn.NatsConn().ConnectedUrl(),
		natsConn.NatsConn().ConnectedServerId(),
	)

	pgxPool, err := postgresql.NewPgxConn(cfg)
	if err != nil {
		appLogger.Fatalf("NewPgxConn: %+v", err)
	}
	appLogger.Infof("PostgreSQL connected: %+v", pgxPool.Stat().TotalConns())

	s := server.NewServer(appLogger, cfg, natsConn, pgxPool, tracer, redisClient)

	appLogger.Fatal(s.Run())
}
