package ingredient_handlers

import (
	"KitchenMistakeErazer/backend/handlers"
	"KitchenMistakeErazer/backend/models"
	"KitchenMistakeErazer/backend/repository"
	"KitchenMistakeErazer/backend/usecases/ingredients"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
)

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
		logrus.Errorf("error while executing inserting ingredient: %s", err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ingredient inserted successfully",
	})
}

func (i *IngredientsHandler) RemoveIngredient(c echo.Context) error {
	id, err := handlers.GetIdFromEndpoint(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err,
		})
	}

	repository := repository.NewIngredientRepository(i.container.DB)
	ingredientRemover := ingredients.NewRemoveIngredient(repository)
	err = ingredientRemover.Execute(id)
	if err != nil {
		logrus.Errorf("error while executing removing ingredient: %s", err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ingredient deleted successfully",
	})
}

func (i *IngredientsHandler) UpdateIngredientHandler(c echo.Context) error {
	id, err := handlers.GetIdFromEndpoint(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err,
		})
	}

	var input models.Ingredient
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err,
		})
	}

	repository := repository.NewIngredientRepository(i.container.DB)
	newUpdateIngredient := ingredients.NewUpdateIngredient(repository)
	err = newUpdateIngredient.Execute(input, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ingredient updated successfully",
	})
}
