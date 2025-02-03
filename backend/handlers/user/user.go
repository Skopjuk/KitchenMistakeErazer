package user

import (
	"KitchenMistakeErazer/backend/container"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	container *container.Container
}

func NewUserHandler(container *container.Container) *UserHandler {
	return &UserHandler{container: container}
}

func (u *UserHandler) SetRoutes(g *echo.Group) {
	g.PUT("/:id", u.UpdateUser)
	g.DELETE("/:id", u.DeleteUser)
	g.PUT("/change_password/:id", u.UpdateUsersPassword)
	g.GET("/:id", u.GetUserByID)
}
