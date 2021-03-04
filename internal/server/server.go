package server

import (
	"github.com/AleksK1NG/nats-streaming/config"
	"github.com/AleksK1NG/nats-streaming/pkg/logger"
)

type server struct {
	log logger.Logger
	cfg *config.Config
}
