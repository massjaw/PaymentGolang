package usecase

import (
	"Merchant-Bank/model/dto/req"
	"Merchant-Bank/repository"
)

type WalletUsecase interface {
	GetWallet(userId int) (string, error)
	TopUp(value *req.TopUp, userId int) (string, error)
}

type walletUsecase struct {
	walletRepo repository.WalletRepo
}

func (u *walletUsecase) GetWallet(userId int) (string, error) {
	return u.walletRepo.GetWallet(userId)
}

func (u *walletUsecase) TopUp(value *req.TopUp, userId int) (string, error) {
	return u.walletRepo.TopUp(value, userId)
}

func NewWalletUsecase(walletRepo repository.WalletRepo) WalletUsecase {
	return &walletUsecase{
		walletRepo: walletRepo,
	}
}
