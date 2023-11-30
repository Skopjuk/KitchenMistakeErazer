package repository

import (
	"KitchenMistakeErazer/backend/models"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type RecipesRepository struct {
	db *sqlx.DB
}

func NewRecipesRepository(db *sqlx.DB) *RecipesRepository {
	return &RecipesRepository{db: db}
}

func (r *RecipesRepository) InsertRecipe(recipe models.Recipe) (err error) {
	query := "INSERT INTO recipes (recipe_name, description, user_id, sourness, saltiness, acidity, sweetness, hot, calories, fat, protein, carbs) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)"
	_, err = r.db.Query(query, recipe.RecipeName, recipe.Description, recipe.UserId, recipe.Sourness, recipe.Saltiness, recipe.Acidity, recipe.Sweetness, recipe.Hot, recipe.Calories, recipe.Fat, recipe.Protein, recipe.Carbs)
	if err != nil {
		logrus.Errorf("error while inserting recipe")
	}

	return err
}
