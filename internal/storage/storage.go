package storage

import "github.com/alekseyprokopov/storage/internal/file"

type Storage struct {
}

func New() *Storage {
	return &Storage{}
}

func (s *Storage) Upload(filename string, blob []byte) (*file.File, error) {
	return file.New(filename, blob)
}
