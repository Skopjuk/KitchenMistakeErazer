package ingredients

import "KitchenMistakeErazer/backend/models"

type GetIngredient struct {
	repository FindIngredient
}

func (g GetIngredient) Execute(id int) (ingredient models.Ingredient, err error) {
	return g.repository.FindIngredient(id)
}
