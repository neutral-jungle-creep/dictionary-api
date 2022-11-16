package storage

import (
	"dictionary/pkg/domain"
	"github.com/jackc/pgx/v4"
)

type Storage struct {
	Dictionary
}

func NewStorage(conn *pgx.Conn) *Storage {
	return &Storage{
		Dictionary: NewWordStorage(conn),
	}
}

type Dictionary interface {
	GetAllFromWords() ([]domain.Word, error)
	GetWordId(*domain.Word) int
	AddWordToDB(word *domain.Word) error
}
