package recipe_version

import "KitchenMistakeErazer/backend/models"

type InsertRecipeToRecipeVersion interface {
	InsertRecipeVersion(recipe models.Recipe) error
}
