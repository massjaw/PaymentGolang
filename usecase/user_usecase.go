package usecase

import (
	"Merchant-Bank/model/entity"
	"Merchant-Bank/repository"
)

type UserUsecase interface {
	Register(newUser *entity.User, confirmPass string) (entity.User, error)
	Login(username string, password string) (string, error)
}

type userUsecase struct {
	userRepo repository.UserRepo
}

func (u *userUsecase) Register(newUser *entity.User, confirmPass string) (entity.User, error) {
	return u.userRepo.CreateUser(newUser, confirmPass)
}

func (u *userUsecase) Login(username string, password string) (string, error) {
	return u.userRepo.UserLogin(username, password)
}

func NewUserUsecase(userRepo repository.UserRepo) UserUsecase {
	return &userUsecase{
		userRepo: userRepo,
	}
}
