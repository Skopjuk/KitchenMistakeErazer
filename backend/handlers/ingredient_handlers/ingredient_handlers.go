package ingredient_handlers

import (
	"KitchenMistakeErazer/backend/models"
	"KitchenMistakeErazer/backend/repository"
	"KitchenMistakeErazer/backend/usecases/ingredients"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
)

type IngreditentsParams struct {
}

//type AddIngredientParams struct {
//	Name              string  `json:"name,omitempty"`
//	Amount            float32 `json:"amount,omitempty"`
//	MeasurementUnitId int     `json:"measurement_unit_id,omitempty"`
//}

func (i *IngredientsHandler) AddIngredient(c echo.Context) error {
	var input models.Ingredient
	if err := c.Bind(&input); err != nil {
		logrus.Errorf("error while binding ingredient: %s", err)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "internal error",
		})
	}

	logrus.Println(input)

	repository := repository.NewIngredientRepository(i.container.DB)
	insertIngredient := ingredients.NewInsertIngredient(repository)

	err := insertIngredient.Execute(input)
	if err != nil {
		logrus.Errorf("error while executing inserting ingredient params: %s", err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ingredient inserted successfully",
	})
}
