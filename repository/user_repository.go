package repository

import (
	"database/sql"
	"log"

	"Merchant-Bank/model/entity"
	"Merchant-Bank/utils"
)

type UserRepo interface {
	CreateUser(newUser *entity.User, confirmPass string) (entity.User, error)
	Login(username string, password string) (string, error)
}

type userRepo struct {
	db *sql.DB
}

func (r *userRepo) CreateUser(newUser *entity.User, confirmPass string) (entity.User, error) {

	query := "INSERT INTO users (user_username,  user_email, user_password) VALUES($1, $2, $3) RETURNING id"

	var userId int
	err := r.db.QueryRow(query, newUser.Username, newUser.Email, newUser.Password).Scan(&userId)
	if err != nil {
		log.Println(err)
		return entity.User{}, err
	}
	newUser.Id = userId
	newUser.Password = ""
	return *newUser, nil
}

func (r *userRepo) Login(username string, password string) (string, error) {
	query := "SELECT id, user_username,  user_email, user_password FROM users WHERE username = $1"

	u := entity.UserLogin{}

	row := r.db.QueryRow(query, username)
	err := row.Scan(&u.UserId, &u.Username, &u.Email, &u.Password)
	if err != nil {
		return "", err
	}

	if u.Password != password {
		return "", err
	}

	token, err := utils.GenerateToken(u.UserId)

	if err != nil {
		return "", err
	}

	return token, nil
}

func NewMemberRepository(db *sql.DB) UserRepo {
	repo := new(userRepo)
	repo.db = db
	return repo
}
