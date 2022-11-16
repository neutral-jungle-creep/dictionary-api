package service

import (
	"dictionary/pkg/domain"
	"dictionary/pkg/service/dto"
	"dictionary/pkg/storage"
)

type EditService struct {
	storage storage.Edit
}

func NewEditService(storage storage.Edit) *EditService {
	return &EditService{
		storage: storage,
	}
}

func (s *EditService) AddWord(wordDto *dto.WordDto) error {
	word := domain.NewWord(wordDto.Id, wordDto.ForeignWord, wordDto.Translation)
	if err := s.storage.AddWordToDB(word); err != nil {
		return err
	}
	return nil
}
