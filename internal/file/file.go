package file

import "github.com/google/uuid"

type File struct {
	ID   uuid.UUID
	Name string
	Data []byte
}

func New(filename string, blob []byte) (*File, error) {
	uid, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	return &File{
		Name: filename,
		Data: blob,
		ID:   uid,
	}, nil
}
