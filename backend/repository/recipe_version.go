package repository

import (
	"KitchenMistakeErazer/backend/models"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type RecipeVersionRepository struct {
	db *sqlx.DB
}

func NewRecipeVersionRepository(db *sqlx.DB) *RecipeVersionRepository {
	return &RecipeVersionRepository{db: db}
}

func (r *RecipeVersionRepository) InsertRecipeVersion(recipe models.Recipe) error {
	query := "INSERT INTO recipe_versions (recipe_name, description, user_id, sourness, saltiness, acidity, sweetness, hot, calories, fat, protein, carbs) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)"
	_, err := r.db.Query(query, recipe.RecipeName, recipe.Description, recipe.UserId, recipe.Sourness, recipe.Saltiness, recipe.Acidity, recipe.Sweetness, recipe.Hot, recipe.Calories, recipe.Fat, recipe.Protein, recipe.Carbs)
	if err != nil {
		logrus.Errorf("error while inserting recipe")
	}

	return err
}
