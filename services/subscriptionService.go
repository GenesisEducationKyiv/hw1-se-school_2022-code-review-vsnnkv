package services

import "net/mail"

type SubscriptionService struct {
	fileService FileService
}

func NewSubscriptionService(f FileService) *SubscriptionService {
	return &SubscriptionService{fileService: f}
}

func (s *SubscriptionService) SaveEmail(email string) int {
	if !isEmailValid(email) {
		return 409
	}
	return s.fileService.repository.SaveEmailToFile(email)
}

func isEmailValid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
