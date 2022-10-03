package services

import (
	"errors"
	"github.com/vsnnkv/btcApplicationGo/tools"
	"net/mail"
)

type SubscriptionServiceInterface interface {
	SaveEmail(email string) error
}

type SubscriptionService struct {
	emailService EmailService
	logger       *tools.LoggerStruct
}

func NewSubscriptionService(f EmailService, l *tools.LoggerStruct) *SubscriptionService {
	return &SubscriptionService{emailService: f, logger: l}
}

func (s *SubscriptionService) SaveEmail(email string) error {
	if !isEmailValid(email) {
		err := errors.New("email not valid")
		s.logger.LogError(err.Error())
		return err
	}
	exist, err := s.emailService.repository.IsExists(email)

	if err != nil {
		err := errors.New("Помилка сервера")
		s.logger.LogError(err.Error())
		return err
	}

	if exist {
		err := errors.New("Email вже було додано")
		s.logger.LogError(err.Error())
		return err
	}

	err = s.emailService.repository.SaveEmailToFile(email)
	if err != nil {
		err := errors.New("Помилка збереження файла")
		s.logger.LogError(err.Error())
		return err
	}
	return nil
}

func isEmailValid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
