package storage

import (
	"github.com/jackc/pgx/v4"
)

type Storage struct {
	Show
}

func NewStorage(conn *pgx.Conn) *Storage {
	return &Storage{
		Show: NewShowStorage(conn),
	}
}

type Show interface {
	ShowAllWords()
}
