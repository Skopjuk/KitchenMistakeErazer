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
	g.POST("/", u.SignUp)
	g.GET("/", u.GetAllUsers)
	g.PUT("/:id", u.UpdateUser)
	g.DELETE("/:id", u.DeleteUser)
	g.PUT("/change_password/:id", u.UpdateUsersPassword)
}
