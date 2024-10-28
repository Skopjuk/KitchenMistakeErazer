package ingredients

import (
	"KitchenMistakeErazer/backend/models"
	"KitchenMistakeErazer/backend/repository"
)

type UpdateIngredient struct {
	repository repository.IngredientRepository
}

func NewUpdateIngredient(repository repository.IngredientRepository) *UpdateIngredient {
	return &UpdateIngredient{repository: repository}
}

func (u UpdateIngredient) Execute(ingredient models.Ingredient, id int) error {
	if ingredient, err := u.repository.FindIngredient(id); ingredient.Name == "" && err != nil {
		return err
	}

	return u.repository.ChangeIngredient(id, ingredient)
}
