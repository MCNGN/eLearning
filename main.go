package main

import (
	"fmt"

	"github.com/MCNGN/eLearning/domain"
	"github.com/MCNGN/eLearning/infrastructure"
	"github.com/MCNGN/eLearning/services"
)

func main() {
	userRepo := infrastructure.NewMemoryUserRepository()
	fileRepo := infrastructure.NewMemoryFileRepository()
	tutorRepo := infrastructure.NewMemoryTutorRepository()
	userService := services.NewUserService(userRepo, fileRepo)
	tutorService := services.NewTutorService(tutorRepo, fileRepo)

	file := &domain.File{
		ID:   "file2",
	}
	
	// Register a new tutor
	_, err := tutorService.RegisterTutor("jane", "doe")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Attempt to login with the registered tutor
	_, err = tutorService.LoginTutor("jane", "doe")
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Attempt to upload a file with the registered tutor
	err = tutorService.UploadFile("jane", file)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
    
	fmt.Println("--------------------------------")
	// Register a new user
    _, err = userService.RegisterUser("john", "doe", )
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

	// Attempt to login with the registered user
	_, err = userService.LoginUser("john", "doe")
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Attempt to download a file with the registered user
	err = userService.DownloadFile("john", "file1")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

}
