package users

import (
	"KitchenMistakeErazer/backend/container"
	"github.com/labstack/echo/v4"
)

type UsersHandler struct {
	container *container.Container
}

func NewUsersHandler(container *container.Container) *UsersHandler {
	return &UsersHandler{container: container}
}

func (u *UsersHandler) SetRoutes(g *echo.Group) {
	g.GET("/", u.GetAllUsers)
}
