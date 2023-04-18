package entity

import "time"

type PaymentDetail struct {
	Id         int       `json:"transaction_id"`
	Payment_id int       `json:"payment_id"`
	userId     int       `json:"-"`
	Amount     float32   `json:"amount"`
	DateTime   time.Time `json:"transaction_datetime"`
}
