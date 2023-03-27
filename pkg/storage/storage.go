package storage

import "github.com/alekseyprokopov/storage/internal/storage"

func NewStorage() *storage.Storage{
	return storage.New()
}