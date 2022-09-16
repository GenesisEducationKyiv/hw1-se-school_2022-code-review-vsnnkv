package services

import (
	"net/mail"
)

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
	code := s.fileService.repository.SaveEmailToFile(email)

	switch code {
	case 200:
		return 200, "Email додано"
	case 400:
		return 400, "Email вже було додано"
	default:
		return 500, "Помилка сервера"
	}
}

func isEmailValid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
