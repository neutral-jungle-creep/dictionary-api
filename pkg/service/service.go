package service

import (
	"dictionary/pkg/domain"
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
	GetAllWords() ([]domain.Word, error)
	AddWord(wordDto *dto.WordDto) error
}
