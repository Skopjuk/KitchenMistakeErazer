package ingredients

import "KitchenMistakeErazer/backend/models"

type AddIngredient interface {
	AddIngredient(ingredient models.Ingredient) error
}
