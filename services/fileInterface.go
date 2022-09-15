package services

type FileRepository interface {
	SaveEmailToFile(email string) int
	//GetEmails() []string
}

type FileService struct {
	repository FileRepository
}

func NewFileService(r FileRepository) *FileService {
	return &FileService{repository: r}
}
