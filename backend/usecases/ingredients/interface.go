package ingredients

import "KitchenMistakeErazer/backend/models"

type AddIngredient interface {
	AddIngredient(ingredient models.Ingredient) error
}

type DeleteIngredient interface {
	DeleteIngredient(id int) error
}

type ChangeIngredient interface {
	ChangeIngredient(id int, ingredient models.Ingredient) error
}

type FindIngredient interface {
	FindIngredient(id int) (ingredient models.Ingredient, err error)
}
