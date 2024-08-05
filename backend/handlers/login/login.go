package login

import (
	"KitchenMistakeErazer/backend/container"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type SignInHandler struct {
	container *container.Container
}

func NewSignInHandler(container *container.Container) *SignInHandler {
	return &SignInHandler{container: container}
}

func (u *SignInHandler) SetRoutes(g *echo.Group) {
	//fileServer := http.FileServer(http.Dir("./static"))
	//
	//g.GET("/", echo.WrapHandler(http.StripPrefix("/", fileServer)))
	g.Use(middleware.Static("/Users/ksenia/KitchenMistakeErazer/static"))
}
