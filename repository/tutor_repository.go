package repository

import (
	"github.com/MCNGN/eLearning/domain"
)

type TutorRepository interface {
	Save(tutor *domain.Tutor) error
	FindByUsername(username string) (*domain.Tutor, error)
}
