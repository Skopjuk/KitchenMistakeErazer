package users

import (
	"KitchenMistakeErazer/backend/models"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type Authentication struct {
	repository GetUserByEmail
}

func NewAuthentication(repository GetUserByEmail) *Authentication {
	return &Authentication{repository: repository}
}

type AuthenticationAttributes struct {
	Email    string
	Password string
}

func (a *Authentication) Execute(attributes AuthenticationAttributes) (user models.User, err error) {
	user, err = a.repository.GetUserByEmail(attributes.Email)
	if err != nil {
		logrus.Error("can not execute usecase: ", err)
		return models.User{}, err
	}

	err = bcrypt.CompareHashAndPassword(user.PasswordDigest, []byte(attributes.Password))
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
