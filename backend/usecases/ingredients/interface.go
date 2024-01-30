package ingredients

type AddIngredient interface {
	AddIngredient(name string, amount float32, measurementUnit string) error
}
