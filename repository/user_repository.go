package repository

import (
	"github.com/MCNGN/eLearning/domain"
)

type UserRepository interface {
	Save(user *domain.User) error
	FindByUsername(username string) (*domain.User, error)
}
