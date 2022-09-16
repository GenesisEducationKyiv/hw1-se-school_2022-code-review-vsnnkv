package services

import (
	"github.com/vsnnkv/btcApplicationGo/config"
	"gopkg.in/gomail.v2"
	"strconv"
)

type NotificationService struct {
	rateService RateService
	fileService FileService
}

func NewNotificationService(r RateService, f FileService) *NotificationService {
	return &NotificationService{rateService: r, fileService: f}
}

func (n *NotificationService) SendEmails() (int, string) {
	emails := n.fileService.repository.GetEmails()

	if len(emails) == 0 {
		return 409, "Відсутні emailʼи"
	}

	rate, err := n.rateService.GetRate()
	if err != nil {
		return 500, "Помилка отримання курсу"
	}

	var cfg = config.Get()
	address := cfg.EmailAddress
	password := cfg.EmailPassword
	host := cfg.SMTPHost
	port, _ := strconv.Atoi(cfg.SMTPPort)

	msg := "курс BTC до UAH відповідно до данних сайту coingecko.com складає: " + strconv.FormatInt(rate, 10)

	m := gomail.NewMessage()
	m.SetHeader("From", address)
	m.SetHeader("To", emails...)
	m.SetHeader("Subject", "Курс BTC до UAH")
	m.SetBody("text/html", msg)

	d := gomail.NewDialer(host, port, address, password)

	if err := d.DialAndSend(m); err != nil {
		panic(err)
		return 500, "Помилка відправки emailʼів"
	}
	return 200, "email відправлено"
}
