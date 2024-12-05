package measurement_units

type RemoveMeasurementUnit struct {
	repository DeleteMeasurementUnit
}

func (r RemoveMeasurementUnit) Execute(id int) error {
	return r.repository.DeleteMeasurementUnit(id)
}
