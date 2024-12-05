package users

import (
	"KitchenMistakeErazer/backend/models"
)

type CreateUserProfile struct {
	repository InsertUser
}

func NewCreateUserProfile(repository InsertUser) *CreateUserProfile {
	return &CreateUserProfile{repository: repository}
}

func (c *CreateUserProfile) Execute(attributes UserAttributes) (id int, err error) {
	err = ParametersValidation(attributes)
	if err != nil {
		return 0, err
	}

	id, err = c.repository.InsertUser(models.User{
		FirstName:      attributes.FirstName,
		LastName:       attributes.LastName,
		Email:          attributes.Email,
		PasswordDigest: PasswordHashing(attributes.Password),
	})

	return id, err
}
