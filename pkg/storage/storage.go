package storage

import (
	"dictionary/pkg/domain"
	"dictionary/pkg/service/dto"
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
	GetAllFromWords() ([]dto.WordDto, error)
	GetWordId(*domain.Word) int
	AddWordToDB(word *domain.Word) error
}
