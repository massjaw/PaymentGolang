package entity

import "time"

type Payment struct {
	Id         int       `json:"transaction_id"`
	SenderId   int       `json:"sender_id"`
	ReceiverId int       `json:"receiver_id"`
	Amount     float32   `json:"amount"`
	DateTime   time.Time `json:"transaction_datetime"`
}
