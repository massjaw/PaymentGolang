package req

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserRegist struct {
	Username        string `json:"username"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"confirm_password"`
}
