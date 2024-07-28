package services

import (
	"fmt"
	"github.com/MCNGN/eLearning/domain"
	"errors"
	"github.com/MCNGN/eLearning/repository"
	"github.com/MCNGN/eLearning/utils"
)

type TutorService struct {
	tutorRepository repository.TutorRepository
	fileRepository repository.FileRepository
}

func NewTutorService(tutorRepository repository.TutorRepository, fileRepository repository.FileRepository) *TutorService {
	return &TutorService{tutorRepository: tutorRepository, fileRepository: fileRepository}
}

func (s *TutorService) RegisterTutor(username string, password string) (*domain.Tutor, error) {
	existingTutor, _ := s.tutorRepository.FindByUsername(username)
	if existingTutor != nil {
		return nil, errors.New("tutor already exists")
	}

	tutor := domain.NewTutor(utils.GenerateID(), username, password)
	err := s.tutorRepository.Save(tutor)
	if err != nil {
		return nil, err
	}

	fmt.Println("Berhasil register tutor")
	return tutor, nil
}

func (s *TutorService) LoginTutor(username string, password string) (*domain.Tutor, error) {
	tutor, err := s.tutorRepository.FindByUsername(username)

	if err != nil {
		return nil, errors.New("tutor not found")
	}

	if tutor.Password != password {
		return nil, errors.New("incorrect password")
	}

	tutor.LoginStatus = true
	fmt.Println("Berhasil login tutor")
	return tutor, nil
}

func (s *TutorService) UploadFile(username string, file *domain.File) error {
	tutor, err := s.tutorRepository.FindByUsername(username)
	if err != nil {
		return errors.New("tutor not found")
	}

	if !tutor.LoginStatus {
		return errors.New("tutor not logged in")
	}

	err = s.fileRepository.Save(file)
	if err != nil {
		return err
	}

	fmt.Println("Berhasil upload file")
	return nil
}