package users

import "KitchenMistakeErazer/backend/models"

type InsertUser interface {
	InsertUser(user models.User) (err error)
}

type UpdateUser interface {
	UpdateUser(user models.User, id int) (err error)
}

type UpdatePassword interface {
	UpdatePassword(password string, id int) (err error)
}

type DeleteUser interface {
	DeleteUser(id int) (err error)
}
