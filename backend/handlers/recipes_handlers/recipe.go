package recipes_handlers

import (
	"KitchenMistakeErazer/backend/container"
	"github.com/labstack/echo/v4"
)

type RecipesHandler struct {
	container *container.Container
}

func NewRecipesHandler(container *container.Container) *RecipesHandler {
	return &RecipesHandler{container: container}
}

func (r *RecipesHandler) SetRoutes(g *echo.Group) {
	g.POST("/", r.AddRecipe)
}
