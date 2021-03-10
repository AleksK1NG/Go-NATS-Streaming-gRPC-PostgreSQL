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

// NewPublisher Nats publisher constructor
func NewPublisher(stanConn stan.Conn) *publisher {
	return &publisher{stanConn: stanConn}
}

// Publish Publish will publish to the cluster and wait for an ACK
func (p *publisher) Publish(subject string, data []byte) error {
	return p.stanConn.Publish(subject, data)
}

// PublishAsync PublishAsync will publish to the cluster and asynchronously process the ACK or error state. It will return the GUID for the message being sent.
func (p *publisher) PublishAsync(subject string, data []byte, ah stan.AckHandler) (string, error) {
	return p.stanConn.PublishAsync(subject, data, ah)
}
