package repository

import (
	"Merchant-Bank/model/dto/req"
	"Merchant-Bank/model/entity"
	"database/sql"
	"errors"
	"log"
	"strconv"
	"time"
)

type PaymentRepo interface {
	CreatePayment(newTransfer *req.Transfer, senderId int) (entity.Payment, error)
}

type paymentRepo struct {
	db *sql.DB
}

func (r *paymentRepo) CreatePayment(newPayment *req.Transfer, senderId int) (entity.Payment, error) {
	log.Println((newPayment))

	tx, err := r.db.Begin()
	if err != nil {
		return entity.Payment{}, err
	}

	var receiptientId string
	query := `SELECT id FROM users WHERE user_username = $1`
	err = tx.QueryRow(query, newPayment.ReceiptUsername).Scan(&receiptientId)
	if err != nil {
		tx.Rollback()
		return entity.Payment{}, err
	}

	receiptId, _ := strconv.Atoi(receiptientId)

	if senderId == receiptId {
		err := errors.New("can't transfer to yourself")
		tx.Rollback()
		return entity.Payment{}, err
	}

	var senderBalance float32
	query = "SELECT balance FROM users_wallet WHERE user_id = $1"
	err = tx.QueryRow(query, senderId).Scan(&senderBalance)
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return entity.Payment{}, err
	} else {
		log.Println("Get SenderBalance")
	}
	if newPayment.TransferAmount > senderBalance {
		err := errors.New("insuficient balance")
		tx.Rollback()
		return entity.Payment{}, err
	}

	// kirim uang
	query = `update users_wallet SET balance = balance - $1 WHERE user_id = $2`
	_, err = tx.Exec(query, newPayment.TransferAmount, senderId)
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return entity.Payment{}, err
	} else {
		log.Println("funds transfered")
	}

	// Terima uang
	query = `update users_wallet SET balance = balance + $1 WHERE user_id = $2`
	_, err = tx.Exec(query, newPayment.TransferAmount, receiptId)
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return entity.Payment{}, err
	} else {
		log.Println("funds received")
	}

	var dateTime time.Time
	var transferId int
	query = `INSERT INTO payment_log (sender_id, receipt_id, amount) VALUES ($1, $2, $3) returning id, date_time`
	err = tx.QueryRow(query, senderId, receiptId, newPayment.TransferAmount).Scan(&transferId, &dateTime)
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return entity.Payment{}, err
	}

	transferedMoney := newPayment.TransferAmount * -1
	query = `INSERT INTO payment_log_detail (payment_id, user_id, amount) VALUES ($1, $2, $3), ($1, $4, $5)`
	_, err = tx.Exec(query, transferId, receiptId, newPayment.TransferAmount, senderId, transferedMoney)
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return entity.Payment{}, err
	}

	err = tx.Commit()
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return entity.Payment{}, err
	} else {
		log.Println("Commited")
	}

	return entity.Payment{
		Id:         transferId,
		SenderId:   senderId,
		ReceiverId: receiptId,
		Amount:     newPayment.TransferAmount,
		DateTime:   dateTime,
	}, nil
}

func NewPaymentRepository(db *sql.DB) PaymentRepo {
	repo := new(paymentRepo)
	repo.db = db
	return repo
}
