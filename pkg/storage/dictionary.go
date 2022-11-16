package storage

import (
	"context"
	"dictionary/pkg/domain"
	"github.com/jackc/pgx/v4"
)

const (
	getAllWords = `SELECT * FROM english_words`
)

type WordStorage struct {
	conn *pgx.Conn
}

func NewWordStorage(conn *pgx.Conn) *WordStorage {
	return &WordStorage{
		conn: conn,
	}
}

func (s *WordStorage) GetAllFromEnglishWords() ([]domain.Word, error) {
	var words []domain.Word
	result, err := s.conn.Query(context.Background(), getAllWords)
	if err != nil {
		return nil, err
	}

	for result.Next() {
		var word domain.Word
		result.Scan(&word.Id, &word.ForeignWord, &word.Translation)
		words = append(words, word)
	}

	return words, nil
}

func (s *WordStorage) GetWordFromDB() {
	//TODO
}

func (s *WordStorage) AddWordToDB(word *domain.Word) error {
	_, err := s.conn.Exec(context.Background(), ``,
		word.ForeignWord,
		word.Translation,
	)
	return err
}
