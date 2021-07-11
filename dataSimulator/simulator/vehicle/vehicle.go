package vehicle

import (
	"github.com/samtholiya/dataSimulator/simulator/vehicle/location"
)

type LocationSensor interface {
	Location() location.Location
}

type Vehicle struct {
	VIN      string
	location LocationSensor
}

func NewSimulator(vin string, lat, long float64, distanceDiff float64) Vehicle {
	simulator := location.NewSimulator(lat, long, distanceDiff)
	return Vehicle{
		VIN:      vin,
		location: &simulator,
	}
}

func (c Vehicle) Location() location.Location {
	return c.location.Location()
}
