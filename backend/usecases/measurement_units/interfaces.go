package measurement_units

type AddMeasurementUnit interface {
	AddMeasurementUnit(unitName string) error
}

type DeleteMeasurementUnit interface {
	DeleteMeasurementUnit(id int) error
}

type FindMeasurementUnit interface {
	FindMeasurementUnit(id int) error
}

//удаляем юнит если он не используется больше ни  каких рецептах
// если мы удалем ингридиент, то мы проверяем есть ли этот рецепт, то мы проверяем есть ли в нем уникальные ингредиенты
// если мы удалем ингредиент то мы проверяем есть ли в нем уникальные единицы измерения
//1. то есть мы берем айди ингредиента и находим все номера межурмент юнитовв с ним связанные
//2. потом берем каждый ингредиент и проверяем есть ли еще записи с ним в таболице ingredient_measurement_unit
