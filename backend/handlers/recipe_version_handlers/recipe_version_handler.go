package recipe_version_handlers

import (
	"KitchenMistakeErazer/backend/handlers"
	"KitchenMistakeErazer/backend/repository"
	"KitchenMistakeErazer/backend/usecases/recipe_version"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (r *RecipesVersionHandler) ShowAllVersionsOfRecipe(c echo.Context) (err error) {
	idInt, err := handlers.GetIdFromEndpoint(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": fmt.Sprintf("users id %d can not be parsed", idInt),
		})
	}

	recipeVersionRepository := repository.NewRecipeVersionRepository(r.container.DB)
	newShowAllRecipesByUserId := recipe_version.NewGetAllVersionsOfRecipe(recipeVersionRepository)
	result, err := newShowAllRecipesByUserId.Execute(idInt)
	if err != nil {
		logrus.Errorf("problem while extracting recipe versions from DB: %s", err)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": fmt.Sprintf("problem while extracting recipe versions from DB: %s", err),
		})
	}

	err = c.JSON(http.StatusOK, map[string]interface{}{
		"recipe_versions": result,
	})
	if err != nil {
		logrus.Errorf("troubles with sending http status: %s", err)
	}

	return err
}
