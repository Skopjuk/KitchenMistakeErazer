package models

type MeasurementUnit struct {
	Id   int    `db:"id"`
	Unit string `db:"unit"`
}
