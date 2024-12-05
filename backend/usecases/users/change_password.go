package users

import "errors"

type ChangePassword struct {
	repository UpdatePassword
}

func NewChangePassword(repository UpdatePassword) *ChangePassword {
	return &ChangePassword{repository: repository}
}

func (c *ChangePassword) Execute(id int, password string) error {
	if len(password) < 6 {
		return errors.New("password is too short")
	}

	hashedPassword := PasswordHashing(password)

	return c.repository.UpdatePassword(hashedPassword, id)
}
