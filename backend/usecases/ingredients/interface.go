package ingredients

import "KitchenMistakeErazer/backend/models"

type AddIngredient interface {
	AddIngredient(ingredient models.Ingredient) error
}

type DeleteIngredient interface {
	DeleteIngredient(id int) error
}
