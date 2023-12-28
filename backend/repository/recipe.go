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

func (r *RecipesRepository) InsertRecipe(recipe models.Recipe) (id int, err error) {
	query := "INSERT INTO recipes (recipe_name, description, user_id, recipe_version_id, sourness, saltiness, acidity, sweetness, hot, calories, fat, protein, carbs) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13) returning id"
	err = r.db.QueryRow(query, recipe.RecipeName, recipe.Description, recipe.UserId, recipe.RecipeVersionId, recipe.Sourness, recipe.Saltiness, recipe.Acidity, recipe.Sweetness, recipe.Hot, recipe.Calories, recipe.Fat, recipe.Protein, recipe.Carbs).Scan(&id)
	if err != nil {
		logrus.Errorf("error while inserting recipe")
	}

	logrus.Info("inserted recipe id: %d", id)

	return id, err
}

func (r *RecipesRepository) ShowAllRecipes(skip, paginationLimit string) (recipes []models.Recipe) {
	query := "SELECT * FROM recipes LIMIT $1 OFFSET $2"
	err := r.db.Select(&recipes, query, paginationLimit, skip)
	if err != nil {
		logrus.Errorf("error while extracting recipes list from db: %s", err)
	}

	return recipes
}

func (r *RecipesRepository) UpdateRecipe(recipe models.Recipe, id int) (err error) {
	query := "UPDATE recipes SET recipe_name=$1, description=$2, user_id=$3, recipe_version_id=$4, sourness=$5, saltiness=$6, acidity=$7, sweetness=$8, hot=$9, calories=$10, fat=$11, protein=$12, carbs=$13 WHERE id = $14"
	_, err = r.db.Query(query, recipe.RecipeName, recipe.Description, recipe.UserId, recipe.RecipeVersionId, recipe.Sourness, recipe.Saltiness, recipe.Acidity, recipe.Sweetness, recipe.Hot, recipe.Calories, recipe.Fat, recipe.Protein, recipe.Carbs, id)
	if err != nil {
		logrus.Errorf("query problem:%s", err)
	}

	return err
}
