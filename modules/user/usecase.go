package user

import (
	"github.com/dibimbing-satkom-indo/onion-architecture-go/entities"
	"github.com/dibimbing-satkom-indo/onion-architecture-go/repositories"
	"time"
)

type UseCase struct {
	userRepo repositories.UserRepo
}

type UseCaseInterface interface {
	CreateUser(user UserParam) (entities.User, error)
	GetUserById(id uint) (entities.User, error)
}

func (uc UseCase) GetUserById(id uint) (entities.User, error) {
	var user, err = uc.userRepo.GetUserById(id)
	return user, err
}

func (uc UseCase) CreateUser(user UserParam) (entities.User, error) {
	var newUser *entities.User

	newUser = &entities.User{
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}

	_, err := uc.userRepo.CreateUser(newUser)
	if err != nil {
		return *newUser, err
	}
	return *newUser, nil
}
