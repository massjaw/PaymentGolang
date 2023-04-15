package entity

import "time"

type Transaction struct {
	Id         int       `json:"transaction_id"`
	SenderId   int       `json:"sender_id"`
	ReceiverId int       `json:"receiver_id"`
	Amount     float32   `json:"amount"`
	DateTime   time.Time `json:"transaction_datetime"`
}

type Transfer struct {
	ReceiptUsername string  `json:"receipt_username"`
	TransferAmount  float32 `json:"transfer_amount"`
	Description     string  `json:"Description"`
}
