package login

import (
	"KitchenMistakeErazer/backend/container"
	"github.com/labstack/echo/v4"
)

type SignInHandler struct {
	container *container.Container
}

func NewSignInHandler(container *container.Container) *SignInHandler {
	return &SignInHandler{container: container}
}

func (u *SignInHandler) SetRoutes(g *echo.Group) {
	g.POST("/", u.SignIn)
}
