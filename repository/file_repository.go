package repository

import "github.com/MCNGN/eLearning/domain"

type FileRepository interface {
    Save(file *domain.File) error
    FindByID(fileID string) (*domain.File, error)
}
