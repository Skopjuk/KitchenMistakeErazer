package recipe_version

import "KitchenMistakeErazer/backend/models"

type GetAllVersionsOfRecipe struct {
	repository ShowAllVersionsOfRecipe
}

func NewGetAllVersionsOfRecipe(repository ShowAllVersionsOfRecipe) *GetAllVersionsOfRecipe {
	return &GetAllVersionsOfRecipe{repository: repository}
}

func (g GetAllVersionsOfRecipe) Execute(id int) (recipeVersionsList []models.RecipeVersion, err error) {
	return g.repository.ShowAllVersionsOfRecipe(id)
}
