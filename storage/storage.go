package storage

import (
	"github.com/jmoiron/sqlx"

	"github.com/JasurbekUz/ToDo-service/storage/postgres"
	"github.com/JasurbekUz/ToDo-service/storage/repo"
)

// IStorage ...
type IStorage interface {
	Todo() repo.TodoStorageI
}

type storagePg struct {
	db       *sqlx.DB
	todoRepo repo.TodoStorageI
}

// NewStoragePg ...
func NewStoragePg(db *sqlx.DB) *storagePg {
	return &storagePg{
		db:       db,
		todoRepo: postgres.NewTodoRepo(db),
	}
}

func (s storagePg) Todo() repo.TodoStorageI {
	return s.todoRepo
}
