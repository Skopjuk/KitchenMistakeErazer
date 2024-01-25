package recipe_version

import "KitchenMistakeErazer/backend/models"

type InsertRecipeToRecipeVersion interface {
	InsertRecipeVersion(recipe models.RecipeVersion, id int, lastRecipeVersion uint) error
}

type GetLatestVersionOfRecipe interface {
	GetLatestVersionOfRecipe(recipeId int) (lastVersion uint, err error)
}
