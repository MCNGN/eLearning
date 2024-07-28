package infrastructure

import (
	"errors"

	"github.com/MCNGN/eLearning/domain"
)

type MemoryFileRepository struct {
	file map[string]*domain.File
}

func NewMemoryFileRepository() *MemoryFileRepository {
	return &MemoryFileRepository{file: make(map[string]*domain.File)}
}

func (r *MemoryFileRepository) Save(file *domain.File) error {
	r.file[file.ID] = file
	return nil
}

func (r *MemoryFileRepository) FindByID(fileID string) (*domain.File, error) {
	for _, file := range r.file {
		if file.ID == fileID {
			return file, nil
		}
	}
	return nil, errors.New("file not found")
}
