package services

import "github.com/vsnnkv/btcApplicationGo/repository"

type FileService struct {
	repository repository.FileRepository
}

func NewFileService(r repository.FileRepository) *FileService {
	return &FileService{repository: r}
}
