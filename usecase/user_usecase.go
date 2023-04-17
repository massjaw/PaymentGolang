package usecase

import (
	"Merchant-Bank/model/dto/req"
	"Merchant-Bank/model/entity"
	"Merchant-Bank/repository"
)

type UserUsecase interface {
	Register(newUser *req.UserRegist) (entity.User, error)
	Login(userInfo *req.UserLogin) (string, error)
}

type userUsecase struct {
	userRepo repository.UserRepo
}

func (u *userUsecase) Register(newUser *req.UserRegist) (entity.User, error) {
	return u.userRepo.CreateUser(newUser)
}

func (u *userUsecase) Login(userInfo *req.UserLogin) (string, error) {
	return u.userRepo.UserLogin(userInfo)
}

func NewUserUsecase(userRepo repository.UserRepo) UserUsecase {
	return &userUsecase{
		userRepo: userRepo,
	}
}
