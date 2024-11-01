package users

import "KitchenMistakeErazer/backend/models"

type InsertUser interface {
	InsertUser(user models.User) (id int, err error)
}

type UpdateUser interface {
	UpdateUser(user models.User, id int) (err error)
}

type UpdatePassword interface {
	UpdatePassword(password []byte, id int) (err error)
}

type DeleteUser interface {
	DeleteUser(id int) (err error)
}

type ShowAllUsers interface {
	ShowAllUsers(skip, paginationLimit string) (usersList []models.User, err error)
}

type GetUserById interface {
	GetUserById(id int) (user models.User, err error)
}

type GetUserByEmail interface {
	GetUserByEmail(email string) (user models.User, err error)
}
