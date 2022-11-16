package service

import (
	"dictionary/pkg/domain"
	"dictionary/pkg/service/dto"
	"dictionary/pkg/storage"
)

type WordService struct {
	storage storage.Dictionary
}

func NewWordService(storage storage.Dictionary) *WordService {
	return &WordService{
		storage: storage,
	}
}

func (s *WordService) GetAllWords() ([]dto.WordDto, error) {
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

func (s *WordService) AddWord(wordDto *dto.WordDto) error {
	word := domain.NewWord(wordDto.Id, wordDto.ForeignWord, wordDto.Translation)
	if err := s.storage.AddWordToDB(word); err != nil {
		return err
	}
	return nil
}
