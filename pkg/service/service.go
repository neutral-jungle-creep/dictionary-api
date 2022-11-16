package service

import (
	"dictionary/pkg/service/dto"
	"dictionary/pkg/storage"
)

type Service struct {
	Dictionary
}

func NewService(storage *storage.Storage) *Service {
	return &Service{
		Dictionary: NewWordService(storage.Dictionary),
	}
}

type Dictionary interface {
	GetAllWords() ([]dto.WordDto, error)
	AddWord(wordDto *dto.WordDto) error
}
