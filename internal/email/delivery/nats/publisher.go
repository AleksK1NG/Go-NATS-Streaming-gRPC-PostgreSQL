package nats

import (
	"github.com/nats-io/stan.go"
)

type Publisher interface {
	Publish(subject string, data []byte) error
	PublishAsync(subject string, data []byte, ah stan.AckHandler) (string, error)
}

type publisher struct {
	stanConn stan.Conn
}

func NewPublisher(stanConn stan.Conn) *publisher {
	return &publisher{stanConn: stanConn}
}

func (p *publisher) Publish(subject string, data []byte) error {
	return p.stanConn.Publish(subject, data)
}

func (p *publisher) PublishAsync(subject string, data []byte, ah stan.AckHandler) (string, error) {
	return p.stanConn.PublishAsync(subject, data, ah)
}
