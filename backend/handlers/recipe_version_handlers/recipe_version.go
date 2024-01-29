package recipe_version_handlers

import (
	"KitchenMistakeErazer/backend/container"
	"github.com/labstack/echo/v4"
)

type RecipesVersionHandler struct {
	container *container.Container
}

func NewRecipesVersionHandler(container *container.Container) *RecipesVersionHandler {
	return &RecipesVersionHandler{container: container}
}

func (r *RecipesVersionHandler) SetRoutes(g *echo.Group) {
	g.GET("/:id", r.ShowAllVersionsOfRecipe)
}
