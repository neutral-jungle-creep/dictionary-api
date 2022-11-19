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

func (s *WordService) GetAllWords() ([]domain.Word, error) {
	var words []domain.Word

	wordsDto, err := s.storage.GetAllFromWords()
	if err != nil {
		return nil, err
	}

	for _, wordDto := range wordsDto {
		words = append(words, *domain.NewWord(wordDto.Id, wordDto.ForeignWord, wordDto.Translation, wordDto.Learned))
	}
	return words, nil
}

func (s *WordService) AddWord(wordDto *dto.WordDto) error {
	word := domain.NewWord(wordDto.Id, wordDto.ForeignWord, wordDto.Translation, wordDto.Learned)

	if err := s.storage.AddWordToDB(word); err != nil {
		return err
	}
	return nil
}
