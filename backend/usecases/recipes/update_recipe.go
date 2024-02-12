package recipes

import (
	"KitchenMistakeErazer/backend/models"
	"KitchenMistakeErazer/backend/repository"
	"errors"
	"github.com/sirupsen/logrus"
)

type ChangeRecipe struct {
	recipeRepository        UpdateRecipe
	recipeVersionRepository repository.RecipeVersionRepository
}

func NewChangeRecipe(recipeRepository UpdateRecipe, recipeVersionRepository repository.RecipeVersionRepository) *ChangeRecipe {
	return &ChangeRecipe{recipeRepository: recipeRepository, recipeVersionRepository: recipeVersionRepository}
}

func (c *ChangeRecipe) Execute(attributes RecipeAttributes, id int) error {
	err := checkIfRecipeAttributesValid(attributes)
	if err != nil {
		return err
	}

	lastRecipeVersion, err := c.recipeVersionRepository.GetLatestVersionOfRecipe(id)
	if err != nil {
		return err
	}

	recipe := models.Recipe{
		UserId: attributes.UserId,
	}

	recipeVersion := models.RecipeVersion{
		RecipeName:  attributes.RecipeName,
		Description: attributes.Description,
		Sourness:    attributes.Sourness,
		Saltiness:   attributes.Saltiness,
		Acidity:     attributes.Acidity,
		Sweetness:   attributes.Sweetness,
		Hot:         attributes.Hot,
		Calories:    attributes.Calories,
		Fat:         attributes.Fat,
		Protein:     attributes.Protein,
		Carbs:       attributes.Carbs,
	}

	err = c.recipeVersionRepository.InsertRecipeVersion(recipeVersion, id, lastRecipeVersion+1)
	if err != nil {
		logrus.Errorf("error while inserting recipe version: %s", err)
	}

	err = c.recipeRepository.UpdateRecipe(recipe, id)
	if err != nil {
		logrus.Errorf("error while updating recipe: %s", err)
	}

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
		return errors.New("sourness can not be more then 6")
	} else if attributes.Saltiness > 6 {
		return errors.New("saltiness can not be more then 6")
	} else if attributes.Acidity > 6 {
		return errors.New("saltiness can not be more then 6")
	} else if attributes.Sweetness > 6 {
		return errors.New("saltiness can not be more then 6")
	} else if attributes.Hot > 6 {
		return errors.New("saltiness can not be more then 6")
	} else if attributes.Calories > 6 {
		return errors.New("saltiness can not be more then 6")
	} else if attributes.Fat > 6 {
		return errors.New("saltiness can not be more then 6")
	} else if attributes.Protein > 6 {
		return errors.New("saltiness can not be more then 6")
	} else if attributes.Carbs > 6 {
		return errors.New("saltiness can not be more then 6")
	}

	return nil
}
