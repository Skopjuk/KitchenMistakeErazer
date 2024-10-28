package measurement_unit_handlers

import (
	"KitchenMistakeErazer/backend/container"
	"github.com/labstack/echo/v4"
)

type MeasurementUnitHandler struct {
	container *container.Container
}

func NewMeasurementUnitHandler(container *container.Container) *MeasurementUnitHandler {
	return &MeasurementUnitHandler{container: container}
}

func (m *MeasurementUnitHandler) SetRoutes(g *echo.Group) {
	g.POST("/", m.InsertMeasurementUnit)
}
