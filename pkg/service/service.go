package service

import (
	"dictionary/pkg/service/dto"
	"dictionary/pkg/storage"
)

type Service struct {
	Show
	Edit
}

func NewService(storage *storage.Storage) *Service {
	return &Service{
		Show: NewShowService(storage.Show),
		Edit: NewEditService(storage.Edit),
	}
}

type Show interface {
	GetAllWords() ([]dto.WordDto, error)
}

type Edit interface {
	AddWord(wordDto *dto.WordDto) error
}
