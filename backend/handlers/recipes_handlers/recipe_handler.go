package recipes_handlers

import (
	"KitchenMistakeErazer/backend/repository"
	"KitchenMistakeErazer/backend/usecases/recipes"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

type AddRecipeParams struct {
	RecipeName  string `json:"recipe_name"`
	Description string `json:"description"`
	UserId      uint   `json:"user_id"`
	Sourness    uint   `json:"sourness"`
	Saltiness   uint   `json:"saltiness"`
	Acidity     uint   `json:"acidity"`
	Sweetness   uint   `json:"sweetness"`
	Hot         uint   `json:"hot"`
	Calories    uint   `json:"calories"`
	Fat         uint   `json:"fat"`
	Protein     uint   `json:"protein"`
	Carbs       uint   `json:"carbs"`
}

func (r *RecipesHandler) AddRecipe(c echo.Context) error {
	var input AddRecipeParams
	if err := c.Bind(&input); err != nil {
		logrus.Errorf("error while binding recipe: %s", err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	logrus.Info(input.RecipeName)

	recipeRepository := repository.NewRecipesRepository(r.container.DB)
	newAddRecipe := recipes.NewCreateRecipe(recipeRepository)
	err := newAddRecipe.Execute(recipes.RecipeAttributes{
		RecipeName:  input.RecipeName,
		Description: input.Description,
		UserId:      input.UserId,
		Sourness:    input.Sourness,
		Saltiness:   input.Saltiness,
		Acidity:     input.Acidity,
		Sweetness:   input.Sweetness,
		Hot:         input.Hot,
		Calories:    input.Calories,
		Fat:         input.Fat,
		Protein:     input.Protein,
		Carbs:       input.Carbs,
	})

	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value") {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"error": "this recipe already exist",
			})
		}

		logrus.Errorf("error while executing inserting recipe params: %s", err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	err = c.JSON(http.StatusOK, map[string]interface{}{
		"added_recipe": input,
	})

	return nil
}
