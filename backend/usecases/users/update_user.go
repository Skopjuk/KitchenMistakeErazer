package users

import (
	"KitchenMistakeErazer/backend/models"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"regexp"
)

type UpdateUserProfile struct {
	repository UpdateUser
}

func NewUpdateUserProfile(repository UpdateUser) *UpdateUserProfile {
	return &UpdateUserProfile{repository: repository}
}

type UpdateUserAttributes struct {
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Email     string `json:"email, omitempty"`
}

func (c *UpdateUserProfile) Execute(attributes UpdateUserAttributes, id int) error {
	err := validateUser(attributes)
	if err != nil {
		logrus.Errorf("error while updating user: %s", err)
		return err
	}

	return c.repository.UpdateUser(models.User{
		FirstName: attributes.FirstName,
		LastName:  attributes.LastName,
		Email:     attributes.Email,
	}, id)
}

func validateUser(attributes UpdateUserAttributes) error {
	emailRegEx := `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`

	if len(attributes.FirstName) < 2 {
		return fmt.Errorf("first name is too short")
	} else if len(attributes.FirstName) > 50 {
		return fmt.Errorf("first name is too long")
	} else if len(attributes.LastName) < 2 {
		return fmt.Errorf("first name is too short")
	} else if len(attributes.LastName) > 50 {
		return fmt.Errorf("last name is too long")
	} else if len(attributes.LastName) > 50 {
		return errors.New("last name is too long")
	} else if !regexp.MustCompile(emailRegEx).MatchString(attributes.Email) {
		return errors.New("email is not valid")
	}

	return nil
}
