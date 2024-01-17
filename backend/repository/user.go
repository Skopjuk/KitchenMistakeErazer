package repository

import (
	"KitchenMistakeErazer/backend/models"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"time"
)

type UsersRepository struct {
	db *sqlx.DB
}

func NewUsersRepository(db *sqlx.DB) *UsersRepository {
	return &UsersRepository{db: db}
}

func (u *UsersRepository) InsertUser(user models.User) error {
	query := "INSERT INTO kitchen_users (first_name, last_name, email, password) values ($1, $2, $3, $4)"
	_, err := u.db.Query(query, user.FirstName, user.LastName, user.Email, user.PasswordDigest)
	if err != nil {
		logrus.Errorf("error while inserting user")
		return err
	}

	return nil
}

func (u *UsersRepository) ShowAllUsers(skip, paginationLimit string) (usersList []models.User, err error) {
	query := "SELECT * FROM kitchen_users LIMIT $1 OFFSET $2"
	err = u.db.Select(&usersList, query, paginationLimit, skip)

	if err != nil {
		logrus.Errorf("error while selecting users from kitchen_users table: %s", err)
	}

	return usersList, err
}

func (u *UsersRepository) UpdateUser(user models.User, id int) (err error) {
	query := "UPDATE kitchen_users SET first_name=$1, last_name=$2, email=$3, updated_at=$4 WHERE id=$5"
	_, err = u.db.Query(query, user.FirstName, user.LastName, user.Email, time.Now(), id)
	if err != nil {
		logrus.Errorf("query problem: %s", err)
	}

	return err
}

func (u *UsersRepository) GetUserById(id int) (user models.User, err error) {
	query := "SELECT * FROM kitchen_users WHERE id=$1"
	err = u.db.Get(&user, query, id)
	if err != nil {
		logrus.Errorf("user with id %d wasn't found, with error: %s", id, err)
	}

	return user, err
}

func (u *UsersRepository) DeleteUser(id int) (err error) {
	query := "DELETE FROM kitchen_users WHERE id=$1"
	_, err = u.db.Query(query, id)
	if err != nil {
		logrus.Errorf("deletion of user with id %d was unseccessful: %s", id, err)
	}

	return err
}

func (u *UsersRepository) UpdatePassword(password []byte, id int) (err error) {
	query := "UPDATE kitchen_users SET password_digest=$1, updated_at=$2 WHERE id=$3"
	_, err = u.db.Query(query, password, time.Now(), id)
	if err != nil {
		logrus.Errorf("password changing was unseccessful: %s", err)
	}

	return err
}
