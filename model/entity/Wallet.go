package entity

type Wallet struct {
	Id      int     `json:"wallet_id"`
	UserId  int     `json:"user_id"`
	Balance float32 `json:"balance"`
}
