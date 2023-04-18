package repository

import (
	"Merchant-Bank/model/dto/req"
	"database/sql"
	"fmt"
	"strconv"
)

type WalletRepo interface {
	GetWallet(userId int) (string, error)
	TopUp(value *req.TopUp, userId int) (string, error)
}

type walletRepo struct {
	db *sql.DB
}

func (r *walletRepo) GetWallet(userId int) (string, error) {

	var balance float32
	query := `SELECT balance FROM users_wallet WHERE user_id = $1`
	err := r.db.QueryRow(query, userId).Scan(&balance)
	if err != nil {
		return "-", err
	}

	balanceConv := strconv.Itoa(int(balance))

	result := "Rp. " + balanceConv
	return result, nil

}

func (r *walletRepo) TopUp(value *req.TopUp, userId int) (string, error) {
	query := `UPDATE users_wallet SET balance = balance + $1 WHERE user_id = $2`
	_, err := r.db.Exec(query, value.Amount, userId)
	if err != nil {
		return "Top-Up Failed", err
	}
	return fmt.Sprintf("top up rp. %v success", value.Amount), nil
}

func NewWalletRepository(db *sql.DB) WalletRepo {
	repo := new(walletRepo)
	repo.db = db
	return repo
}
