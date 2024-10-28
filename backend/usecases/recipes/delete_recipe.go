package recipes

import "KitchenMistakeErazer/backend/usecases/recipe_version"

type RemoveRecipe struct {
	repository          DeleteRecipe
	versionRepository   recipe_version.DeleteRecipeVersion
	getRecipeRepository GetRecipe
}

func NewRemoveRecipe(repository DeleteRecipe, versionRepository recipe_version.DeleteRecipeVersion, getRecipeRepository GetRecipe) *RemoveRecipe {
	return &RemoveRecipe{repository: repository, versionRepository: versionRepository, getRecipeRepository: getRecipeRepository}
}

func (r *RemoveRecipe) Execute(id int) error {
	_, err := r.getRecipeRepository.GetRecipe(id)
	if err != nil {
		return err
	}

	err = r.versionRepository.DeleteRecipeVersion(id)
	if err != nil {
		return err
	}

	err = r.repository.DeleteRecipe(id)

	return err
}
