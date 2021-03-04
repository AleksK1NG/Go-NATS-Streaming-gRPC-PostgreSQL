package v1

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	successRequests = promauto.NewCounter(prometheus.CounterOpts{
		Name: "http_email_success_incoming_messages_total",
		Help: "The total number of success incoming email HTTP requests",
	})
	errorRequests = promauto.NewCounter(prometheus.CounterOpts{
		Name: "http_email_error_incoming_message_total",
		Help: "The total number of error incoming email HTTP requests",
	})
	createRequests = promauto.NewCounter(prometheus.CounterOpts{
		Name: "http_email_create_incoming_requests_total",
		Help: "The total number of incoming create email HTTP requests",
	})
	updateRequests = promauto.NewCounter(prometheus.CounterOpts{
		Name: "http_email_update_incoming_requests_total",
		Help: "The total number of incoming update email HTTP requests",
	})
	getByIdRequests = promauto.NewCounter(prometheus.CounterOpts{
		Name: "http_email_get_by_id_incoming_requests_total",
		Help: "The total number of incoming get by id email HTTP requests",
	})
	searchRequests = promauto.NewCounter(prometheus.CounterOpts{
		Name: "http_email_search_incoming_requests_total",
		Help: "The total number of incoming search email HTTP requests",
	})
)
