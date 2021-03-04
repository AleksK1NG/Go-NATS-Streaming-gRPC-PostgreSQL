package nats

import (
	"sync"
	"time"

	"github.com/AleksK1NG/nats-streaming/internal/email"
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

func NewEmailSubscriber(stanConn stan.Conn, log logger.Logger, emailUC email.UseCase, validator *validator.Validate) *emailSubscriber {
	return &emailSubscriber{stanConn: stanConn, log: log, emailUC: emailUC, validator: validator}
}

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

func (s *emailSubscriber) Run() {
	go s.Subscribe(createEmailSubject, emailGroupName, createEmailWorkers, s.createEmail)
}

func (s *emailSubscriber) createEmail(msg *stan.Msg) {
	s.log.Infof("create:email message: %+v", msg)
	if err := msg.Ack(); err != nil {
		s.log.Errorf("msg.Ack: %+v", err)
	}
}
