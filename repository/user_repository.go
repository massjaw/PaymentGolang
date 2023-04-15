package repository

import (
	"database/sql"
	"log"

	"Merchant-Bank/model/entity"
)

type UserRepo interface {
	CreateUser(newUser *entity.User, confirmPass string) (entity.User, error)
	Login(username string, password string) (entity.User, error)
}

type userRepo struct {
	db *sql.DB
}

func (r *userRepo) CreateUser(newUser *entity.User, confirmPass string) (entity.User, error) {

	query := "INSERT INTO users (user_username,  user_email, user_password) VALUES($1, $2, $3) RETURNING id"

	var userId int
	err := r.db.QueryRow(query, newUser.Username, newUser.Password, newUser.Email).Scan(&userId)
	if err != nil {
		log.Println(err)
		return entity.User{}, err
	}
	newUser.Password = ""
	return *newUser, nil
}

func (r *userRepo) Login(username string, password string) (entity.User, error) {
	query := "SELCET user_username,  user_email, user_password FROM users WHERE username = $1"

	var err error
	u := entity.UserLogin{}

	row := r.db.QueryRow(query, username)
	err = row.Scan(&u.UserId)

	return entity.User{}, nil
}

func NewMemberRepository(db *sql.DB) UserRepo {
	repo := new(userRepo)
	repo.db = db
	return repo
}
