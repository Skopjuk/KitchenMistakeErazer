package recipes

import (
	"KitchenMistakeErazer/backend/models"
	recipe_version "KitchenMistakeErazer/backend/usecases/recipe_version"
	"errors"
)

type CreateRecipe struct {
	repository              InsertRecipe
	repositoryRecipeVersion recipe_version.InsertRecipeToRecipeVersion
}

func NewCreateRecipe(repository InsertRecipe, repositoryRecipeVersion recipe_version.InsertRecipeToRecipeVersion) *CreateRecipe {
	return &CreateRecipe{repository: repository, repositoryRecipeVersion: repositoryRecipeVersion}
}

type RecipeAttributes struct {
	RecipeName      string
	Description     string
	UserId          uint
	RecipeVersionId uint
	Sourness        uint
	Saltiness       uint
	Acidity         uint
	Sweetness       uint
	Hot             uint
	Calories        uint
	Fat             uint
	Protein         uint
	Carbs           uint
}

func (c *CreateRecipe) Execute(attributes RecipeAttributes) (err error) {
	err = CheckIfRecipeAttributesValid(attributes)
	if err != nil {
		return err
	}

	recipe := models.Recipe{
		UserId: attributes.UserId,
	}

	recipeVersion := models.RecipeVersion{
		RecipeName:          attributes.RecipeName,
		Description:         attributes.Description,
		RecipeVersionNumber: 1,
		Sourness:            attributes.Sourness,
		Saltiness:           attributes.Saltiness,
		Acidity:             attributes.Acidity,
		Sweetness:           attributes.Sweetness,
		Hot:                 attributes.Hot,
		Calories:            attributes.Calories,
		Fat:                 attributes.Fat,
		Protein:             attributes.Protein,
		Carbs:               attributes.Carbs,
	}

	id, err := c.repository.InsertRecipe(int(recipe.UserId))
	if err != nil {
		return err
	}

	return c.repositoryRecipeVersion.InsertRecipeVersion(recipeVersion, id, 1)
}

func CheckIfRecipeAttributesValid(attributes RecipeAttributes) (err error) {
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
