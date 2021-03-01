package nats

import (
	"time"

	"github.com/AleksK1NG/nats-streaming/config"
	"github.com/nats-io/stan.go"
)

const (
	connectWait = time.Second * 30
	pubAckWait  = time.Second * 30

	NATS_CLIENT_ID = "NATS_CLIENT_ID"
	NATS_URL       = "NATS_URL"

	CLUSTER_ID = "CLUSTER_ID"
)

func NewNatsConnect(cfg *config.Config) (stan.Conn, error) {
	return stan.Connect(
		cfg.Nats.ClusterID,
		cfg.Nats.ClientID,
		stan.ConnectWait(connectWait),
		stan.PubAckWait(pubAckWait),
		stan.NatsURL(cfg.Nats.URL),
	)
}
