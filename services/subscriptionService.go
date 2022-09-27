package services

import (
	"net/mail"
)

type SubscriptionServiceInterface interface {
	SaveEmail(email string) (int, string)
}

type SubscriptionService struct {
	fileService FileService
}

func NewSubscriptionService(f FileService) *SubscriptionService {
	return &SubscriptionService{fileService: f}
}

func (s *SubscriptionService) SaveEmail(email string) (int, string) {
	if !isEmailValid(email) {
		return 409, "email not valid"
	}
	exist, err := s.fileService.repository.IsExists(email)

	if err != nil {
		return 500, "Помилка сервера"
	}

	if exist {
		return 400, "Email вже було додано"
	}

	err = s.fileService.repository.SaveEmailToFile(email)
	if err != nil {
		return 400, "Помилка збереження файла"
	}
	return 200, "Email додано"
}

func isEmailValid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
