package storage

import (
	"fmt"
	"github.com/alekseyprokopov/storage/internal/file"
	"github.com/google/uuid"
)

type Storage struct {
	Files map[uuid.UUID]*file.File
}

func New() *Storage {
	return &Storage{
		Files: make(map[uuid.UUID]*file.File),
	}
}

func (s *Storage) Upload(filename string, blob []byte) (*file.File, error) {
	File, err := file.New(filename, blob)

	if err != nil {
		return nil, err
	}
	s.Files[File.ID] = File
	return File, err
}

func (s *Storage) GetById(ID uuid.UUID) (*file.File, error) {
	File, ok := s.Files[ID]
	if !ok {
		return nil, fmt.Errorf("can't find file (incorrect ID: %v)", ID)
	}
	return File, nil
}
