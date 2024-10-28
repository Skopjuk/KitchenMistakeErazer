package models

type Ingredient struct {
	Id                int     `db:"id" json:"id,omitempty"`
	Name              string  `db:"name" json:"name,omitempty"`
	Amount            float32 `db:"amount" json:"amount,omitempty"`
	MeasurementUnitId int     `db:"measurement_unit_id" json:"measurement_unit_id,omitempty"`
}
