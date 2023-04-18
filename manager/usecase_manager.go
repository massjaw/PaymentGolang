package manager

import (
	"Merchant-Bank/usecase"
)

type UsecaseManager interface {
	UserUsecase() usecase.UserUsecase
	PaymentUsecase() usecase.PaymentUsecase
	WalletUsecase() usecase.WalletUsecase
}

type usecaseManager struct {
	repoManager RepoManager
}

func (u *usecaseManager) UserUsecase() usecase.UserUsecase {
	return usecase.NewUserUsecase(u.repoManager.UserRepo())
}

func (u *usecaseManager) PaymentUsecase() usecase.PaymentUsecase {
	return usecase.NewPaymentUsecase(u.repoManager.PaymentRepo())
}

func (u *usecaseManager) WalletUsecase() usecase.WalletUsecase {
	return usecase.NewWalletUsecase(u.repoManager.WalletRepo())
}

func NewUsecaseManager(rm RepoManager) UsecaseManager {
	return &usecaseManager{
		repoManager: rm,
	}
}
