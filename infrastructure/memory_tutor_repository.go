package infrastructure

import (
	// "fmt"

	"github.com/MCNGN/eLearning/domain"
	"errors"
)

type MemoryTutorRepository struct {
	tutor map[string]*domain.Tutor
}

func (r *MemoryTutorRepository) Save(tutor *domain.Tutor) error {
	r.tutor[tutor.ID] = tutor
	return nil
}

func (r *MemoryTutorRepository) FindByUsername(username string) (*domain.Tutor, error) {
	for _, user := range r.tutor {
		if user.Name == username {
			return user, nil
		}
	}
	return nil, errors.New("user not found")
}

func NewMemoryTutorRepository() *MemoryTutorRepository {
	return &MemoryTutorRepository{tutor: make(map[string]*domain.Tutor)}
}
