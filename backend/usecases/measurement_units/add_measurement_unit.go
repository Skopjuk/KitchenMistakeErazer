package measurement_units

type InsertUnit struct {
	repository AddMeasurementUnit
}

func NewInsertUnit(repository AddMeasurementUnit) *InsertUnit {
	return &InsertUnit{repository: repository}
}

func (i InsertUnit) Execute(unitName string) error {
	return i.repository.AddMeasurementUnit(unitName)
}
