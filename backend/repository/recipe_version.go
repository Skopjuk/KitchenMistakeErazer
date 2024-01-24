package repository

import (
	"KitchenMistakeErazer/backend/models"
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"time"
)

type RecipeVersionRepository struct {
	db *sqlx.DB
}

func NewRecipeVersionRepository(db *sqlx.DB) *RecipeVersionRepository {
	return &RecipeVersionRepository{db: db}
}

func (r *RecipeVersionRepository) InsertRecipeVersion(recipe models.RecipeVersion, id int, lastRecipeVersion uint) error {
	query := "INSERT INTO recipe_versions (recipe_name, description, recipe_id, recipe_version_id, sourness, saltiness, acidity, sweetness, hot, calories, fat, protein, carbs, created_at) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)"
	_, err := r.db.Query(query, recipe.RecipeName, recipe.Description, id, lastRecipeVersion, recipe.Sourness, recipe.Saltiness, recipe.Acidity, recipe.Sweetness, recipe.Hot, recipe.Calories, recipe.Fat, recipe.Protein, recipe.Carbs, time.Now())
	if err != nil {
		logrus.Errorf("error while inserting recipe version: %s", err)
	}

	return err
}

func (r *RecipeVersionRepository) GetLatestVersionOfRecipe(recipeId int) (lastVersion uint, err error) {
	var recipesV []models.RecipeVersion
	var lastRecipe models.RecipeVersion
	query := "SELECT * FROM recipe_versions WHERE recipe_id=$1"
	err = r.db.Select(&recipesV, query, recipeId)
	if err != nil {
		logrus.Errorf("error while selecting recipes versions from database")
		return 0, err
	}
	if len(recipesV) < 1 {
		logrus.Errorf("no versions in recipe version table")
		return 0, errors.New("no versions in recipe version table")
	}

	lastRecipe = recipesV[len(recipesV)-1]
	lastVersion = lastRecipe.RecipeVersionId

	return lastVersion, err
}
