package ingredients

import "KitchenMistakeErazer/backend/models"

type InsertIngredient struct {
	repository AddIngredient
}

func NewInsertIngredient(repository AddIngredient) *InsertIngredient {
	return &InsertIngredient{repository: repository}
}

func (i InsertIngredient) Execute(ingredient models.Ingredient) error {
	return i.repository.AddIngredient(ingredient)
}
