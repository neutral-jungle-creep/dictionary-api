package storage

import (
	"dictionary/pkg/domain"
	"github.com/jackc/pgx/v4"
)

type Storage struct {
	Show
	Edit
}

func NewStorage(conn *pgx.Conn) *Storage {
	return &Storage{
		Show: NewShowStorage(conn),
		Edit: NewEditStorage(conn),
	}
}

type Show interface {
	GetAllFromEnglishWords() ([]domain.Word, error)
}

type Edit interface {
	AddWordToDB(word *domain.Word) error
}
