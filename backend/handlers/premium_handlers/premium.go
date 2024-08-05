package premium_handlers

import (
	"KitchenMistakeErazer/backend/container"
	"github.com/labstack/echo/v4"
)

type PremiumHandler struct {
	container *container.Container
}

func NewPremiumHandler(container *container.Container) *PremiumHandler {
	return &PremiumHandler{container: container}
}

func (u *PremiumHandler) SetRoutes(g *echo.Group) {
	g.GET("/", u.PremiumHandler)
}
