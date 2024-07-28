package services

import (
	"errors"
	"fmt"
	"github.com/MCNGN/eLearning/domain"
	"github.com/MCNGN/eLearning/repository"
	"github.com/MCNGN/eLearning/utils"
)

type UserService struct {
	userRepository repository.UserRepository
	fileRepository repository.FileRepository
}

func NewUserService(userRepository repository.UserRepository, fileRepository repository.FileRepository) *UserService {
	return &UserService{userRepository: userRepository, fileRepository: fileRepository}
}

func (s *UserService) RegisterUser(username string, password string) (*domain.User, error) {
	existingUser, _ := s.userRepository.FindByUsername(username)
	if existingUser != nil {
		return nil, errors.New("username already taken")
	}

	user := domain.NewUser(utils.GenerateID(), username, password)
	err := s.userRepository.Save(user)
	if err != nil {
		return nil, err
	}

	fmt.Println("Berhasil register user")
	return user, nil
}

func (s *UserService) LoginUser(username string, password string) (*domain.User, error) {
    user, err := s.userRepository.FindByUsername(username)
	
    if err != nil {
        return nil, errors.New("user not found")
    }

    if user.Password != password {
        return nil, errors.New("incorrect password")
    }

	user.LoginStatus = true
    fmt.Println("Berhasil login user")
    return user, nil
}

func (s *UserService) DownloadFile(username string, fileId string) error {
	user, err := s.userRepository.FindByUsername(username)
	if err != nil {
		return errors.New("user not found")
	}

	if !user.LoginStatus {
		return errors.New("user not logged in")
	}

	_, err = s.fileRepository.FindByID(fileId)
	if err != nil {
		return errors.New("file not found")
	} 
	
	fmt.Println("Berhasil download file")
	return nil
}

