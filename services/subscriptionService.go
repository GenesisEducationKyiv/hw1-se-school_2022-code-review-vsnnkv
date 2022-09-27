package services

import (
	"errors"
	"net/mail"
)

type SubscriptionServiceInterface interface {
	SaveEmail(email string) error
}

type SubscriptionService struct {
	fileService FileService
}

func NewSubscriptionService(f FileService) *SubscriptionService {
	return &SubscriptionService{fileService: f}
}

func (s *SubscriptionService) SaveEmail(email string) error {
	if !isEmailValid(email) {
		return errors.New("email not valid")
	}
	exist, err := s.fileService.repository.IsExists(email)

	if err != nil {
		return errors.New("Помилка сервера")
	}

	if exist {
		return errors.New("Email вже було додано")
	}

	err = s.fileService.repository.SaveEmailToFile(email)
	if err != nil {
		return errors.New("Помилка збереження файла")
	}
	return nil
}

func isEmailValid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
