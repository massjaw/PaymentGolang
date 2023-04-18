package usecase

import (
	"Merchant-Bank/model/dto/req"
	"Merchant-Bank/model/entity"
	"Merchant-Bank/repository"
)

type PaymentUsecase interface {
	PaymentProcces(newTransfer *req.Transfer, senderId int) (entity.Payment, error)
}

type paymentUsecase struct {
	paymentRepo repository.PaymentRepo
}

func (u *paymentUsecase) PaymentProcces(newTransfer *req.Transfer, senderId int) (entity.Payment, error) {
	return u.paymentRepo.CreatePayment(newTransfer, senderId)
}

func NewPaymentUsecase(paymentRepo repository.PaymentRepo) PaymentUsecase {
	return &paymentUsecase{
		paymentRepo: paymentRepo,
	}
}
