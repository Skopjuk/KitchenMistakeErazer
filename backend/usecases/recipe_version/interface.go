package recipe_version

import "KitchenMistakeErazer/backend/models"

type InsertRecipeToRecipeVersion interface {
	InsertRecipeVersion(recipe models.RecipeVersion, id int, lastRecipeVersion uint) error
}

type GetLatestVersionOfRecipe interface {
	GetLatestVersionOfRecipe(recipeId int) (lastVersion uint, err error)
}

type DeleteRecipeVersion interface {
	DeleteRecipeVersion(id int) error
}

type ShowAllVersionsOfRecipe interface {
	ShowAllVersionsOfRecipe(id int) (allRecipeVersions []models.RecipeVersion, err error)
}
