package services

import (
	"errors"
	"net/mail"
)

type SubscriptionServiceInterface interface {
	SaveEmail(email string) error
}

type SubscriptionService struct {
	emailService EmailService
}

func NewSubscriptionService(f EmailService) *SubscriptionService {
	return &SubscriptionService{emailService: f}
}

func (s *SubscriptionService) SaveEmail(email string) error {
	if !isEmailValid(email) {
		return errors.New("email not valid")
	}
	exist, err := s.emailService.repository.IsExists(email)

	if err != nil {
		return errors.New("Помилка сервера")
	}

	if exist {
		return errors.New("Email вже було додано")
	}

	err = s.emailService.repository.SaveEmailToFile(email)
	if err != nil {
		return errors.New("Помилка збереження файла")
	}
	return nil
}

func isEmailValid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
