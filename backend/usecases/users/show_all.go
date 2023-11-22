package users

import (
	"KitchenMistakeErazer/backend/models"
	"github.com/sirupsen/logrus"
)

type ShowUsers struct {
	repository ShowAllUsers
}

func NewShowUsers(repository ShowAllUsers) *ShowUsers {
	return &ShowUsers{repository: repository}
}

func (s *ShowUsers) Execute(skip, paginationLimit string) (users []models.User, err error) {
	users, err = s.repository.ShowAllUsers(skip, paginationLimit)
	if err != nil {
		logrus.Errorf("users wern't found: %s", err)
		return nil, err
	}
	return users, err
}
