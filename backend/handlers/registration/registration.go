package registration

import (
	"KitchenMistakeErazer/backend/container"
	"github.com/labstack/echo/v4"
)

type SignUpHandler struct {
	container *container.Container
}

func NewSignUpHandler(container *container.Container) *SignUpHandler {
	return &SignUpHandler{container: container}
}

func (u *SignUpHandler) SetRoutes(g *echo.Group) {
	g.POST("/", u.SignUp)
}
