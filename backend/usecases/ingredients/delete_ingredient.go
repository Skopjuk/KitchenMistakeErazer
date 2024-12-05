package ingredients

type RemoveIngredient struct {
	repository DeleteIngredient
}

func NewRemoveIngredient(repository DeleteIngredient) *RemoveIngredient {
	return &RemoveIngredient{repository: repository}
}

func (r *RemoveIngredient) Execute(id int) error {
	return r.repository.DeleteIngredient(id)
}
