package services

import (
	"errors"
	"github.com/vsnnkv/btcApplicationGo/config"
	"github.com/vsnnkv/btcApplicationGo/tools"
	"gopkg.in/gomail.v2"
	"strconv"
)

type NotificationServiceInterface interface {
	SendEmails() error
}

type NotificationService struct {
	rateService  RateService
	emailService EmailService
	logger       *tools.LoggerStruct
}

func NewNotificationService(r RateService, f EmailService, l *tools.LoggerStruct) *NotificationService {
	return &NotificationService{rateService: r, emailService: f, logger: l}
}

func (n *NotificationService) SendEmails() error {
	emails := n.emailService.repository.GetEmails()

	if len(emails) == 0 {
		err := errors.New("Відсутні emailʼи")
		n.logger.LogError(err.Error())
		return err
	}

	rate, err := n.rateService.GetRate()
	if err != nil {
		err := errors.New("Помилка отримання курсу")
		n.logger.LogError(err.Error())
		return err
	}

	var cfg = config.Get()
	address := cfg.EmailAddress
	password := cfg.EmailPassword
	host := cfg.SMTPHost
	port, _ := strconv.Atoi(cfg.SMTPPort)

	msg := "курс BTC до UAH складає: " + strconv.FormatInt(rate, 10)

	m := gomail.NewMessage()
	m.SetHeader("From", address)
	m.SetHeader("To", emails...)
	m.SetHeader("Subject", "Курс BTC до UAH")
	m.SetBody("text/html", msg)

	d := gomail.NewDialer(host, port, address, password)

	if err := d.DialAndSend(m); err != nil {
		err := errors.New("Помилка відправки emailʼів")
		n.logger.LogError(err.Error())
		return err
	}

	return nil
}
