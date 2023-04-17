package repository

import (
	"database/sql"
	"log"

	"Merchant-Bank/model/dto/req"
	"Merchant-Bank/model/entity"
	"Merchant-Bank/utils"
)

type UserRepo interface {
	CreateUser(newUser *req.UserRegist) (entity.User, error)
	UserLogin(userInfo *req.UserLogin) (string, error)
}

type userRepo struct {
	db *sql.DB
}

func (r *userRepo) CreateUser(newUser *req.UserRegist) (entity.User, error) {

	query := "INSERT INTO users (user_username,  user_email, user_password) VALUES($1, $2, $3) RETURNING id"

	var userId int
	var user entity.User
	err := r.db.QueryRow(query, newUser.Username, newUser.Email, newUser.Password).Scan(&userId)
	if err != nil {
		log.Println(err)
		return entity.User{}, err
	}
	user.Id = userId
	user.Username = newUser.Username
	user.Email = newUser.Email
	user.Password = ""
	return *&user, nil
}

func (r *userRepo) UserLogin(userInfo *req.UserLogin) (string, error) {
	query := "SELECT id, user_username,  user_email, user_password FROM users WHERE user_username = $1"

	u := entity.User{}

	row := r.db.QueryRow(query, userInfo.Username)
	err := row.Scan(&u.Id, &u.Username, &u.Email, &u.Password)
	if err != nil {
		return "", err
	}

	if u.Password != userInfo.Password {
		return "", err
	}

	token, err := utils.GenerateToken(u.Id)

	if err != nil {
		return "", err
	}

	return token, nil
}

func NewUserRepository(db *sql.DB) UserRepo {
	repo := new(userRepo)
	repo.db = db
	return repo
}
