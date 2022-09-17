package services

type FileRepository interface {
	SaveEmailToFile(email string) error
	IsExists(email string) (bool, error)
	GetEmails() []string
}

type FileService struct {
	repository FileRepository
}

func NewFileService(r FileRepository) *FileService {
	return &FileService{repository: r}
}
