package storage

import (
	"context"
	"dictionary/pkg/service/dto"
	"github.com/jackc/pgx/v4"
)

const (
	getAllWords = `SELECT * FROM english_words`
)

type ShowStorage struct {
	conn *pgx.Conn
}

func NewShowStorage(conn *pgx.Conn) *ShowStorage {
	return &ShowStorage{
		conn: conn,
	}
}

func (s *ShowStorage) GetAllFromEnglishWords() ([]dto.WordDto, error) {
	var words []dto.WordDto
	result, err := s.conn.Query(context.Background(), getAllWords)
	if err != nil {
		return nil, err
	}

	for result.Next() {
		var word dto.WordDto
		result.Scan(&word.Id, &word.ForeignWord, &word.Translation)
		words = append(words, word)
	}

	return words, nil
}
