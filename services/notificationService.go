package services

import (
	"errors"
	"github.com/vsnnkv/btcApplicationGo/config"
	"gopkg.in/gomail.v2"
	"strconv"
)

type NotificationServiceInterface interface {
	SendEmails() error
}

type NotificationService struct {
	rateService  RateService
	emailService EmailService
}

func NewNotificationService(r RateService, f EmailService) *NotificationService {
	return &NotificationService{rateService: r, emailService: f}
}

func (n *NotificationService) SendEmails() error {
	emails := n.emailService.repository.GetEmails()

	if len(emails) == 0 {
		return errors.New("Відсутні emailʼи")
	}

	rate, err := n.rateService.GetRate()
	if err != nil {
		return errors.New("Помилка отримання курсу")
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
		panic(err)
		return errors.New("Помилка відправки emailʼів")
	}
	return nil
}
