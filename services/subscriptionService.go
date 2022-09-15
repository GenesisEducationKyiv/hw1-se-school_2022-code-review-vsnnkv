package services

type SubscriptionService struct {
	fileService FileService
}

func NewSubscriptionService(f FileService) *SubscriptionService {
	return &SubscriptionService{fileService: f}
}

func (s *SubscriptionService) SaveEmail(email string) int {

	return s.fileService.repository.SaveEmailToFile(email)
}
