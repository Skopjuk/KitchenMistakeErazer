package server

import (
	"KitchenMistakeErazer/backend/container"
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
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
