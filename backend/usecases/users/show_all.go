package users

import (
	"KitchenMistakeErazer/backend/models"
)

type ShowUsers struct {
	repository ShowAllUsers
}

func NewShowUsers(repository ShowAllUsers) *ShowUsers {
	return &ShowUsers{repository: repository}
}

func (s *ShowUsers) Execute(skip, paginationLimit string) (users []models.User, err error) {
	return s.repository.ShowAllUsers(skip, paginationLimit)
}
