package ingredient_handlers

import (
	"KitchenMistakeErazer/backend/container"
	"github.com/labstack/echo/v4"
)

type IngredientsHandler struct {
	container *container.Container
}

func NewIngredientsHandler(container *container.Container) *IngredientsHandler {
	return &IngredientsHandler{container: container}
}

func (i *IngredientsHandler) SetRoutes(g *echo.Group) {
	g.POST("/", i.AddIngredient)
	g.DELETE("/:id", i.RemoveIngredient)
	g.PATCH("/:id", i.UpdateIngredientHandler)
}
