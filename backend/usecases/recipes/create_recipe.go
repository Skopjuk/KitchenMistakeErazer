package recipes

import (
	"KitchenMistakeErazer/backend/models"
	"errors"
)

type CreateRecipe struct {
	repository InsertRecipe
}

func NewCreateRecipe(repository InsertRecipe) *CreateRecipe {
	return &CreateRecipe{repository: repository}
}

type RecipeAttributes struct {
	RecipeName  string
	Description string
	UserId      uint
	Sourness    uint
	Saltiness   uint
	Acidity     uint
	Sweetness   uint
	Hot         uint
	Calories    uint
	Fat         uint
	Protein     uint
	Carbs       uint
}

func (c *CreateRecipe) Execute(attributes RecipeAttributes) (err error) {
	err = checkIfRecipeAttributesValid(attributes)
	if err != nil {
		return err
	}

	err = c.repository.InsertRecipe(models.Recipe{
		RecipeName:  attributes.RecipeName,
		Description: attributes.Description,
		UserId:      attributes.UserId,
		Sourness:    attributes.Sourness,
		Saltiness:   attributes.Saltiness,
		Acidity:     attributes.Acidity,
		Sweetness:   attributes.Sweetness,
		Hot:         attributes.Hot,
		Calories:    attributes.Calories,
		Fat:         attributes.Fat,
		Protein:     attributes.Protein,
		Carbs:       attributes.Carbs,
	})

	return err
}

func checkIfRecipeAttributesValid(attributes RecipeAttributes) (err error) {
	if len(attributes.RecipeName) < 2 {
		return errors.New("recipe name should be at least 2 symbols")
	} else if len(attributes.RecipeName) > 250 {
		return errors.New("recipe name should be maximum 250 symbols")
	} else if len(attributes.Description) < 100 {
		return errors.New("recipe description should be at least 100 symbols")
	} else if attributes.Sourness > 6 {
		errors.New("sourness can not be more then 6")
	} else if attributes.Saltiness > 6 {
		errors.New("saltiness can not be more then 6")
	} else if attributes.Acidity > 6 {
		errors.New("saltiness can not be more then 6")
	} else if attributes.Sweetness > 6 {
		errors.New("saltiness can not be more then 6")
	} else if attributes.Hot > 6 {
		errors.New("saltiness can not be more then 6")
	} else if attributes.Calories > 6 {
		errors.New("saltiness can not be more then 6")
	} else if attributes.Fat > 6 {
		errors.New("saltiness can not be more then 6")
	} else if attributes.Protein > 6 {
		errors.New("saltiness can not be more then 6")
	} else if attributes.Carbs > 6 {
		errors.New("saltiness can not be more then 6")
	}

	return nil
}
