package email

import (
	"time"

	"github.com/AleksK1NG/nats-streaming/config"
	"github.com/AleksK1NG/nats-streaming/internal/models"
	mail "github.com/xhit/go-simple-mail/v2"
)

type SMTPClient interface {
	SendMail(mail *models.MailData) error
}

type smtpClient struct {
	cfg *config.Config
}

func NewSmtpClient(cfg *config.Config) *smtpClient {
	return &smtpClient{cfg: cfg}
}

// NewEmailSMTPClient connect to mail server and returns SMTP client
func (s *smtpClient) getConn() (*mail.SMTPClient, error) {
	server := mail.NewSMTPClient()

	// SMTP Server
	server.Host = s.cfg.MailService.Host
	server.Port = s.cfg.MailService.Port
	server.Username = s.cfg.MailService.Username
	server.Password = s.cfg.MailService.Password
	server.Encryption = mail.EncryptionNone
	server.ConnectTimeout = s.cfg.MailService.ConnectTimeout * time.Second
	server.SendTimeout = s.cfg.MailService.SendTimeout * time.Second
	server.KeepAlive = s.cfg.MailService.KeepAlive

	return server.Connect()
}

func (s *smtpClient) SendMail(mailData *models.MailData) error {
	conn, err := s.getConn()
	if err != nil {
		return err
	}
	defer conn.Close()

	msg := mail.NewMSG()
	msg.SetFrom(mailData.From)
	msg.AddTo(mailData.To)
	msg.SetSubject(mailData.Subject)
	msg.SetBody(mail.TextPlain, mailData.Content)

	return msg.Send(conn)
}
