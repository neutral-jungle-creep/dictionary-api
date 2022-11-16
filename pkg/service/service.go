package service

import (
	"dictionary/pkg/domain"
	"dictionary/pkg/storage"
)

type Service struct {
	Show
}

func NewService(storage *storage.Storage) *Service {
	return &Service{
		Show: NewShowService(storage.Show),
	}
}

type Show interface {
	GetAllWords() ([]domain.Word, error)
}
