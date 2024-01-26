package recipes

import "KitchenMistakeErazer/backend/usecases/recipe_version"

type RemoveRecipe struct {
	repository                   DeleteRecipe
	versionRepository            recipe_version.DeleteRecipeVersion
	checkIfRecipeExistRepository CheckIfRecipeExist
}

func NewRemoveRecipe(repository DeleteRecipe, versionRepository recipe_version.DeleteRecipeVersion, checkIfRecipeExistRepository CheckIfRecipeExist) *RemoveRecipe {
	return &RemoveRecipe{repository: repository, versionRepository: versionRepository, checkIfRecipeExistRepository: checkIfRecipeExistRepository}
}

func (r *RemoveRecipe) Execute(id int) error {
	err := r.checkIfRecipeExistRepository.CheckIfRecipeExist(id)
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
