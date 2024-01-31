package measurement_unit_handlers

import (
	"KitchenMistakeErazer/backend/models"
	"KitchenMistakeErazer/backend/repository"
	"KitchenMistakeErazer/backend/usecases/measurement_units"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

func (m *MeasurementUnitHandler) InsertMeasurementUnit(c echo.Context) (err error) {
	var input models.MeasurementUnit
	if err := c.Bind(&input); err != nil {
		logrus.Errorf("error while binding input: %s", err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	repository := repository.NewMeasurementUnitsRepository(m.container.DB)
	insertUnit := measurement_units.NewInsertUnit(repository)

	err = insertUnit.Execute(input.Unit)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value") {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"error": "this unit already exist",
			})
		}

		logrus.Errorf("error while executing inserting measurement params: %s", err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	err = c.JSON(http.StatusOK, map[string]interface{}{
		"added_unit": input.Unit,
	})
	if err != nil {
		logrus.Errorf("troubles with sending http status: %s", err)
	}

	return err
}
