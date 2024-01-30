package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type MeasurementUnitsRepository struct {
	db *sqlx.DB
}

func NewMeasurementUnitsRepository(db *sqlx.DB) *MeasurementUnitsRepository {
	return &MeasurementUnitsRepository{db: db}
}

func (m *MeasurementUnitsRepository) AddMeasurementUnit(unitName string) error {
	query := "INSERT INTO measurements (unit) values ($1)"
	_, err := m.db.Query(query, unitName)
	if err != nil {
		logrus.Errorf("error while inserting measurement unit to DB")
	}

	return err
}
