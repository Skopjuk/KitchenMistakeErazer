package ingredients

import "KitchenMistakeErazer/backend/models"

type UpdateIngredient struct {
	repository ChangeIngredient
}

func NewUpdateIngredient(repository ChangeIngredient) *UpdateIngredient {
	return &UpdateIngredient{repository: repository}
}

func (u UpdateIngredient) Execute(ingredient models.Ingredient, id int) error {
	return u.repository.ChangeIngredient(id, ingredient)
}
