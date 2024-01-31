package server

import (
	"KitchenMistakeErazer/backend/container"
	"KitchenMistakeErazer/backend/handlers/ingredient_handlers"
	"KitchenMistakeErazer/backend/handlers/measurement_unit_handlers"
	"KitchenMistakeErazer/backend/handlers/recipe_version_handlers"
	"KitchenMistakeErazer/backend/handlers/recipes_handlers"
	"KitchenMistakeErazer/backend/handlers/users"
	"context"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

func Run(port string, containerInstance container.Container) error {
	e := echo.New()

	routes(e, containerInstance)

	return e.Start(":" + port)
}

func routes(e *echo.Echo, container container.Container) {
	users.NewUsersHandler(&container).SetRoutes(e.Group("/user"))
	recipes_handlers.NewRecipesHandler(&container).SetRoutes(e.Group("/recipe"))
	recipe_version_handlers.NewRecipesVersionHandler(&container).SetRoutes(e.Group("/recipe_versions"))
	measurement_unit_handlers.NewMeasurementUnitHandler(&container).SetRoutes(e.Group("/units"))
	ingredient_handlers.NewIngredientsHandler(&container).SetRoutes(e.Group("/ingredient"))
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
