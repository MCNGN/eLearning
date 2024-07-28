package infrastructure

import (
	// "fmt"

	"github.com/MCNGN/eLearning/domain"
	// "github.com/MCNGN/eLearning/repository"
	"errors"
)

type MemoryUserRepository struct {
	users map[string]*domain.User
}

func (r *MemoryUserRepository) Save(user *domain.User) error {
	r.users[user.ID] = user
	return nil
}

func (r *MemoryUserRepository) FindByUsername(username string) (*domain.User, error) {
	for _, user := range r.users {
		if user.Name == username {
			return user, nil
		}
	}
	return nil, errors.New("user not found")
}

func NewMemoryUserRepository() *MemoryUserRepository {
	return &MemoryUserRepository{users: make(map[string]*domain.User)}
}
