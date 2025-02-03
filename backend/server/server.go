package server

import (
	"KitchenMistakeErazer/backend/container"
	"KitchenMistakeErazer/backend/handlers"
	"KitchenMistakeErazer/backend/handlers/ingredient_handlers"
	"KitchenMistakeErazer/backend/handlers/login"
	"KitchenMistakeErazer/backend/handlers/measurement_unit_handlers"
	"KitchenMistakeErazer/backend/handlers/recipe_version_handlers"
	"KitchenMistakeErazer/backend/handlers/recipes_handlers"
	"KitchenMistakeErazer/backend/handlers/registration"
	"KitchenMistakeErazer/backend/handlers/user"
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
	users.NewUsersHandler(&container).SetRoutes(e.Group("/users", handlers.UserIdentityMiddleware))
	user.NewUserHandler(&container).SetRoutes(e.Group("/user", handlers.UserIdentityMiddleware))
	recipes_handlers.NewRecipesHandler(&container).SetRoutes(e.Group("/recipe"))
	recipe_version_handlers.NewRecipesVersionHandler(&container).SetRoutes(e.Group("/recipe_versions"))
	measurement_unit_handlers.NewMeasurementUnitHandler(&container).SetRoutes(e.Group("/units"))
	ingredient_handlers.NewIngredientsHandler(&container).SetRoutes(e.Group("/ingredient"))
	login.NewSignInHandler(&container).SetRoutes(e.Group("/login"))
	registration.NewSignUpHandler(&container).SetRoutes(e.Group("/sign_up"))
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
