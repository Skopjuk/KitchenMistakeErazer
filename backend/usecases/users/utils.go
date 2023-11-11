package users

import (
	"errors"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"net/mail"
)

type UserAttributes struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
}

func ParametersValidation(attributes UserAttributes) (err error) {
	if len(attributes.FirstName) < 2 {
		return errors.New("first name is too short")
	} else if len(attributes.FirstName) > 50 {
		return errors.New("first name is too long")
	} else if len(attributes.LastName) < 2 {
		return errors.New("last name is too short")
	} else if len(attributes.LastName) > 50 {
		return errors.New("last name is too long")
	} else if _, err = mail.ParseAddress(attributes.Email); err != nil {
		return errors.New("email is not valid")
	} else if len(attributes.Password) < 6 {
		return errors.New("password is too short")
	}
	return nil
}

func PasswordHashing(password string) []byte {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	if err != nil {
		logrus.Error("Password wasn't hashed")
	}
	return hashedPassword
}
