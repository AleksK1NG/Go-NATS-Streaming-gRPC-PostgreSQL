package nats

import (
	"context"
	"encoding/json"
	"sync"
	"time"

	"github.com/AleksK1NG/nats-streaming/internal/email"
	"github.com/AleksK1NG/nats-streaming/internal/models"
	"github.com/AleksK1NG/nats-streaming/pkg/logger"
	"github.com/go-playground/validator/v10"
	"github.com/nats-io/stan.go"
)

type emailSubscriber struct {
	stanConn  stan.Conn
	log       logger.Logger
	emailUC   email.UseCase
	validator *validator.Validate
}

// NewEmailSubscriber email subscriber constructor
func NewEmailSubscriber(stanConn stan.Conn, log logger.Logger, emailUC email.UseCase, validator *validator.Validate) *emailSubscriber {
	return &emailSubscriber{stanConn: stanConn, log: log, emailUC: emailUC, validator: validator}
}

// Subscribe subscribe to subject and run workers with given callback for handling messages
func (s *emailSubscriber) Subscribe(subject, qgroup string, workersNum int, cb stan.MsgHandler) {
	s.log.Infof("Subscribing to Subject: %v, group: %v", subject, qgroup)
	wg := &sync.WaitGroup{}

	for i := 0; i <= workersNum; i++ {
		wg.Add(1)
		go s.runWorker(
			wg,
			i,
			s.stanConn,
			subject,
			qgroup,
			cb,
			stan.SetManualAckMode(),
			stan.AckWait(60*time.Second),
			stan.DurableName("microservice-dur"),
			stan.MaxInflight(25),
		)
	}
	wg.Wait()
}

func (s *emailSubscriber) runWorker(
	wg *sync.WaitGroup,
	workerID int,
	conn stan.Conn,
	subject string,
	qgroup string,
	cb stan.MsgHandler,
	opts ...stan.SubscriptionOption,
) {
	defer wg.Done()

	s.log.Infof("Subscribing worker: %v, subject: %v, qgroup: %v", workerID, subject, qgroup)
	_, err := conn.QueueSubscribe(subject, qgroup, cb, opts...)
	if err != nil {
		s.log.Errorf("Worker: %v, QueueSubscribe: %v", workerID, err)
		if err := conn.Close(); err != nil {
			s.log.Errorf("Worker: %v, conn.Close: %v", workerID, err)
		}
	}
}

// Run start subscribers
func (s *emailSubscriber) Run() {
	go s.Subscribe(createEmailSubject, emailGroupName, createEmailWorkers, s.createEmail)
	go s.Subscribe(sendEmailSubject, emailGroupName, sendEmailWorkers, s.sendEmail)
}

func (s *emailSubscriber) createEmail(msg *stan.Msg) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancelFunc()
	s.log.Infof("createEmail: %+v", msg)
	totalSubscribeMessages.Inc()

	var m models.Email
	if err := json.Unmarshal(msg.Data, &m); err != nil {
		errorSubscribeMessages.Inc()
		s.log.Errorf("json.Unmarshal : %v", err)
		return
	}

	if err := s.emailUC.Create(ctx, &m); err != nil {
		errorSubscribeMessages.Inc()
		s.log.Errorf("emailUC.Create : %v", err)
		return
	}

	if err := msg.Ack(); err != nil {
		s.log.Errorf("msg.Ack: %+v", err)
	}
	successSubscribeMessages.Inc()
}

func (s *emailSubscriber) sendEmail(msg *stan.Msg) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancelFunc()
	s.log.Infof("sendEmail: %+v", msg)
	totalSubscribeMessages.Inc()

	var m models.Email
	if err := json.Unmarshal(msg.Data, &m); err != nil {
		errorSubscribeMessages.Inc()
		s.log.Errorf("json.Unmarshal : %v", err)
		return
	}

	if err := s.emailUC.SendEmail(ctx, &m); err != nil {
		errorSubscribeMessages.Inc()
		s.log.Errorf("emailUC.SendEmail : %v", err)
		return
	}

	if err := msg.Ack(); err != nil {
		s.log.Errorf("msg.Ack: %+v", err)
	}
	successSubscribeMessages.Inc()
}
