package recipes_handlers

import (
	"KitchenMistakeErazer/backend/repository"
	"KitchenMistakeErazer/backend/usecases/recipes"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"strings"
)

var (
	paginationLimit = "10"
)

type RecipeParams struct {
	RecipeName      string `json:"recipe_name"`
	Description     string `json:"description"`
	UserId          uint   `json:"user_id"`
	RecipeVersionId uint   `json:"recipe_version_number"`
	Sourness        uint   `json:"sourness"`
	Saltiness       uint   `json:"saltiness"`
	Acidity         uint   `json:"acidity"`
	Sweetness       uint   `json:"sweetness"`
	Hot             uint   `json:"hot"`
	Calories        uint   `json:"calories"`
	Fat             uint   `json:"fat"`
	Protein         uint   `json:"protein"`
	Carbs           uint   `json:"carbs"`
}

func (r *RecipesHandler) AddRecipe(c echo.Context) error {
	var input RecipeParams
	if err := c.Bind(&input); err != nil {
		logrus.Errorf("error while binding recipe: %s", err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	logrus.Info(input.RecipeName)

	recipeRepository := repository.NewRecipesRepository(r.container.DB)
	recipeVersionRepository := repository.NewRecipeVersionRepository(r.container.DB)
	newAddRecipe := recipes.NewCreateRecipe(recipeRepository, recipeVersionRepository)
	err := newAddRecipe.Execute(recipes.RecipeAttributes{
		RecipeName:      input.RecipeName,
		Description:     input.Description,
		UserId:          input.UserId,
		RecipeVersionId: input.RecipeVersionId,
		Sourness:        input.Sourness,
		Saltiness:       input.Saltiness,
		Acidity:         input.Acidity,
		Sweetness:       input.Sweetness,
		Hot:             input.Hot,
		Calories:        input.Calories,
		Fat:             input.Fat,
		Protein:         input.Protein,
		Carbs:           input.Carbs,
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

func (r *RecipesHandler) GetAllRecipes(c echo.Context) error {
	pageNum := c.QueryParam("page")
	if pageNum == "" {
		pageNum = "1"
	}

	page, err := strconv.Atoi(pageNum)
	if err != nil {
		logrus.Errorf("error while converting page num to int: %s", err)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "error while parsing url",
		})
	}

	skip := strconv.Itoa((page - 1) * 10)
	logrus.Info("attempt to get users list from db")

	recipesRepository := repository.NewRecipesRepository(r.container.DB)
	getAll := recipes.NewGetAllRecipes(recipesRepository)

	recipes := getAll.Execute(skip, paginationLimit)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"recipes": recipes,
	})
}

func (r *RecipesHandler) UpdateRecipe(c echo.Context) error {
	idInt, err := getIdFromEndpoint(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": fmt.Sprintf("users id %d can not be parsed", idInt),
		})
	}

	var input RecipeParams
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "recipe wasn't parsed successfully",
		})
	}

	recipeRepository := repository.NewRecipesRepository(r.container.DB)
	recipeVersionRepository := repository.NewRecipeVersionRepository(r.container.DB)
	newUpdateRecipe := recipes.NewChangeRecipe(recipeRepository, *recipeVersionRepository)
	err = newUpdateRecipe.Execute(recipes.RecipeAttributes{
		RecipeName:      input.RecipeName,
		Description:     input.Description,
		UserId:          input.UserId,
		RecipeVersionId: input.RecipeVersionId,
		Sourness:        input.Sourness,
		Saltiness:       input.Saltiness,
		Acidity:         input.Acidity,
		Sweetness:       input.Sweetness,
		Hot:             input.Hot,
		Calories:        input.Calories,
		Fat:             input.Fat,
		Protein:         input.Protein,
		Carbs:           input.Carbs,
	}, idInt)

	if err != nil {
		logrus.Errorf("error while executing updating recipe params: %s", err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	err = c.JSON(http.StatusOK, map[string]interface{}{
		"updated_recipe": input,
	})

	return nil
}

func (r *RecipesHandler) DeleteRecipe(c echo.Context) error {
	idInt, err := getIdFromEndpoint(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "user id was not parsed from link",
		})
	}

	recipeRepository := repository.NewRecipesRepository(r.container.DB)
	recipeVersionRepository := repository.NewRecipeVersionRepository(r.container.DB)
	removeRecipe := recipes.NewRemoveRecipe(recipeRepository, recipeVersionRepository, recipeRepository)
	err = removeRecipe.Execute(idInt)
	if err != nil {
		logrus.Errorf("recipe with id %d was not removed successfully", idInt)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "user recipe was not delete successfully",
		})
	}

	logrus.Infof("recipe with id %d deleted successfully", idInt)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "recipe delete successfully",
	})
}

func getIdFromEndpoint(c echo.Context) (int, error) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		logrus.Errorf("error of converting id to int. id: %s", id)
		return 0, c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": fmt.Sprintf("id %d can not be parsed", idInt),
		})
	}

	return idInt, nil
}
