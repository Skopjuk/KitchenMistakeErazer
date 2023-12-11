package recipes

import "KitchenMistakeErazer/backend/models"

type GetAllRecipes struct {
	repository ShowAllRecipes
}

func NewGetAllRecipes(repository ShowAllRecipes) *GetAllRecipes {
	return &GetAllRecipes{repository: repository}
}

func (g *GetAllRecipes) Execute(skip, paginationLimit string) []models.Recipe {
	return g.repository.ShowAllRecipes(skip, paginationLimit)
}
