package services

type FileRepository interface {
	SaveEmailToFile(email string) error
	//GetEmails() []string
}

type FileService struct {
	repository FileRepository
}

func NewFileRepository(r FileRepository) *FileService {
	return &FileService{repository: r}
}
