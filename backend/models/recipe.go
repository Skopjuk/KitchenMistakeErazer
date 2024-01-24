package models

import "time"

type Recipe struct {
	Id        int       `db:"id"`
	UserId    uint      `db:"user_id"`
	CreatedAt time.Time `db:"created_at"`
}

type RecipeVersion struct {
	Id              int       `db:"id"`
	RecipeName      string    `db:"recipe_name"`
	Description     string    `db:"description"`
	RecipeId        uint      `db:"recipe_id"`
	RecipeVersionId uint      `db:"recipe_version_id"`
	Sourness        uint      `db:"sourness"`
	Saltiness       uint      `db:"saltiness"`
	Acidity         uint      `db:"acidity"`
	Sweetness       uint      `db:"sweetness"`
	Hot             uint      `db:"hot"`
	Calories        uint      `db:"calories"`
	Fat             uint      `db:"fat"`
	Protein         uint      `db:"protein"`
	Carbs           uint      `db:"carbs"`
	CreatedAt       time.Time `db:"created_at"`
	UpdatedAt       time.Time `db:"updated_at"`
}
