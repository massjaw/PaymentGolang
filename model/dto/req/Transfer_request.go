package req

type Transfer struct {
	ReceiptUsername string  `json:"receipt_username"`
	TransferAmount  float32 `json:"transfer_amount"`
}
