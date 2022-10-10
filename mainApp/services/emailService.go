package services

import (
	"github.com/vsnnkv/btcApplicationGo/infrastructure/repository"
)

type EmailService struct {
	repository repository.EmailRepositoryInterface
}

func NewEmailService(r repository.EmailRepositoryInterface) *EmailService {
	return &EmailService{repository: r}
}
