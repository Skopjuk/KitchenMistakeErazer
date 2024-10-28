package repository

import (
	"KitchenMistakeErazer/backend/models"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type IngredientRepository struct {
	db *sqlx.DB
}

func NewIngredientRepository(db *sqlx.DB) *IngredientRepository {
	return &IngredientRepository{db: db}
}

func (i IngredientRepository) AddIngredient(ingredient models.Ingredient) error {
	query := "INSERT INTO ingredients (name, amount, measurement_unit_id) values ($1, $2, $3)"
	_, err := i.db.Query(query, ingredient.Name, ingredient.Amount, ingredient.MeasurementUnitId)
	if err != nil {
		logrus.Errorf("error while inserting ingredient to DB: %s", err)
	}

	return err
}

func (i IngredientRepository) DeleteIngredient(id int) error {
	query := "DELETE FROM ingredients WHERE id=$1"
	_, err := i.db.Query(query, id)
	if err != nil {
		logrus.Errorf("error while deleting ingredient to DB: %s", err)
	}

	return err
}

func (i IngredientRepository) ChangeIngredient(id int, ingredient models.Ingredient) error {
	query := "UPDATE ingredients SET name=$1, amount=$2, measurement_unit_id=$3 WHERE id=$4"
	_, err := i.db.Query(query, ingredient.Name, ingredient.Amount, ingredient.MeasurementUnitId, id)
	if err != nil {
		logrus.Errorf("error whilt changing ingredient in DB:%s", err)
	}

	return err
}

func (i IngredientRepository) FindIngredient(id int) (ingredient models.Ingredient, err error) {
	query := "SELECT * FROM ingredients WHERE id=$1"
	err = i.db.Get(&ingredient, query, id)
	if err != nil {
		logrus.Errorf("error while looking for ingredient in DB:%s", err)
	}

	return ingredient, err
}
