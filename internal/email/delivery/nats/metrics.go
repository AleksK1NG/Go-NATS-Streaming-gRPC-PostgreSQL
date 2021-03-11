package nats

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	totalSubscribeMessages = promauto.NewCounter(prometheus.CounterOpts{
		Name: "nats_email_incoming_messages_total",
		Help: "The total number of incoming email NATS messages",
	})
	successSubscribeMessages = promauto.NewCounter(prometheus.CounterOpts{
		Name: "nats_email_success_incoming_messages_total",
		Help: "The total number of success email NATS messages",
	})
	errorSubscribeMessages = promauto.NewCounter(prometheus.CounterOpts{
		Name: "nats_email_error_incoming_messages_total",
		Help: "The total number of error email NATS messages",
	})
)
