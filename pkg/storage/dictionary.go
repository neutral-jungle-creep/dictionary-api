package storage

import (
	"context"
	"dictionary/pkg/domain"
	"dictionary/pkg/service/dto"
	"github.com/jackc/pgx/v4"
)

const (
	getAllWords = `SELECT * FROM english_words`
	getWordId   = `SELECT id FROM english_words WHERE word=$1`
	addWord     = `INSERT INTO english_words (word, transl, learned) VALUES ($1, $2, $3)`
)

type WordStorage struct {
	conn *pgx.Conn
}

func NewWordStorage(conn *pgx.Conn) *WordStorage {
	return &WordStorage{
		conn: conn,
	}
}

func (s *WordStorage) GetAllFromWords() ([]dto.WordDto, error) {
	var words []dto.WordDto

	result, err := s.conn.Query(context.Background(), getAllWords)
	if err != nil {
		return nil, err
	}

	for result.Next() {
		var word dto.WordDto
		result.Scan(&word.Id, &word.ForeignWord, &word.Translation, &word.Learned)
		words = append(words, word)
	}

	return words, nil
}

func (s *WordStorage) GetWordId(word *domain.Word) int {
	var wordId int

	result := s.conn.QueryRow(context.Background(), getWordId, word.ForeignWord)
	if err := result.Scan(&wordId); err != nil {
		return 0
	}
	return wordId
}

func (s *WordStorage) AddWordToDB(word *domain.Word) error {
	_, err := s.conn.Exec(context.Background(), addWord,
		word.ForeignWord,
		word.Translation,
		word.Learned,
	)
	return err
}
