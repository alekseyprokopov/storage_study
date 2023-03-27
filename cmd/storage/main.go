package main

import (
	"github.com/alekseyprokopov/storage/internal/storage"
	"log"
)

func main() {
	s := storage.New()

	file, _ := s.Upload("test.txt", []byte("hello"))
	log.Println("uploaded file: ", file)

	data, _ := s.GetById(file.ID)

	log.Println("file from ID: ", data)

}
