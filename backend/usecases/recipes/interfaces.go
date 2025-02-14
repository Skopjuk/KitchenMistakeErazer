package recipes

import "KitchenMistakeErazer/backend/models"

type InsertRecipe interface {
	InsertRecipe(userId int) (id int, err error)
}

type UpdateRecipe interface {
	UpdateRecipe(recipe models.Recipe, id int) (err error)
}

type DeleteRecipe interface {
	DeleteRecipe(id int) (err error)
}

type ShowRecipeByUserId interface {
	ShowRecipeByUserId(userId int) (recipe models.Recipe)
}

type ShowAllRecipes interface {
	ShowAllRecipes(skip, paginationLimit string) (recipes []models.Recipe)
}

type GetRecipe interface {
	GetRecipe(id int) (recipe models.Recipe, err error)
}
