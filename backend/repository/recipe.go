package repository

import (
	"KitchenMistakeErazer/backend/models"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"time"
)

type RecipesRepository struct {
	db *sqlx.DB
}

func NewRecipesRepository(db *sqlx.DB) *RecipesRepository {
	return &RecipesRepository{db: db}
}

func (r *RecipesRepository) InsertRecipe(userId int) (id int, err error) {
	query := "INSERT INTO recipes (user_id, created_at) values ($1, $2) returning id"
	err = r.db.QueryRow(query, userId, time.Now()).Scan(&id)
	if err != nil {
		logrus.Errorf("error while inserting recipe")
	}

	logrus.Infof("inserted recipe id: %d", id)

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
	query := "UPDATE recipes SET user_id=$1 WHERE id = $2"
	_, err = r.db.Query(query, recipe.UserId, id)
	if err != nil {
		logrus.Errorf("query problem:%s", err)
	}

	return err
}

func (r *RecipesRepository) DeleteRecipe(id int) (err error) {
	query := "DELETE FROM recipes WHERE id=$1"
	_, err = r.db.Query(query, id)
	if err != nil {
		logrus.Errorf("delete recipe query problem: %s", err)
	}

	return err
}

func (r *RecipesRepository) GetRecipe(id int) (recipe models.Recipe, err error) {
	query := "SELECT * FROM recipes WHERE id=$1"
	err = r.db.Get(&recipe, query, id)
	if err != nil {
		logrus.Errorf("recipe %d not found", id)
		return recipe, err
	}

	if err != nil {
		logrus.Errorf("checking recipe exictance error: %s", err)
	}

	return recipe, err
}
