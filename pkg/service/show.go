package service

import (
	"dictionary/pkg/service/dto"
	"dictionary/pkg/storage"
)

type ShowService struct {
	storage storage.Show
}

func NewShowService(storage storage.Show) *ShowService {
	return &ShowService{
		storage: storage,
	}
}

func (s *ShowService) GetAllWords() ([]dto.WordDto, error) {
	var wordsDto []dto.WordDto

	words, err := s.storage.GetAllFromEnglishWords()
	if err != nil {
		return nil, err
	}

	for _, word := range words {
		wordsDto = append(wordsDto, *dto.NewWordDto(word.Id, word.ForeignWord, word.Translation))
	}
	return wordsDto, nil
}
