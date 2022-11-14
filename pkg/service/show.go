package service

import (
	"dictionary/pkg/domain"
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

func (s *ShowService) ShowWords() ([]domain.Word, error) {
	var words []domain.Word

	wordsDto, err := s.storage.GetAllWords()
	if err != nil {
		return nil, err
	}

	for _, wordDto := range wordsDto {
		words = append(words, *domain.NewWord(wordDto.Id, wordDto.Word, wordDto.Translation))
	}
	return words, nil
}
