package user

import (
	"KitchenMistakeErazer/backend/models"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type UsersRepository struct {
	db *sqlx.DB
}

func NewUsersRepository(db *sqlx.DB) *UsersRepository {
	return &UsersRepository{db: db}
}

func (u *UsersRepository) InsertUser(user models.User) error {
	query := "INSERT INTO kitchen_users (first_name, last_name, email, password) values ($1, $2, $3, $4)"
	_, err := u.db.Query(query, user.FirstName, user.LastName, user.Email, user.Password)
	if err != nil {
		logrus.Errorf("error while inserting user")
		return err
	}

	return nil
}
