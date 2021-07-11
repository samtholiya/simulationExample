package location

import "math"

const (
	radiusOfEarth   = 6378.1
	brng            = 1.57 // Bearing is 90 degrees converted to radians.
	radianConvertor = math.Pi / 180
	degreeConvertor = 180 / math.Pi
)

type Simulator struct {
	location Location
	distance float64
}

func NewSimulator(lat, long, distance float64) Simulator {
	return Simulator{
		location: Location{Latitude: lat, Longitude: long},
		distance: distance,
	}
}

func (s *Simulator) Location() Location {
	lat1 := (s.location.Latitude) * radianConvertor  // Current lat point converted to radians
	lon1 := (s.location.Longitude) * radianConvertor //Current long point converted to radians
	degree := s.distance / radiusOfEarth
	lat2 := math.Asin(math.Sin(lat1)*math.Cos(degree) +
		math.Cos(lat1)*math.Sin(degree)*math.Cos(brng))

	lon2 := lon1 + math.Atan2(math.Sin(brng)*math.Sin(degree)*math.Cos(lat1),
		math.Cos(degree)-math.Sin(lat1)*math.Sin(lat2))

	lat2 = (lat2 * degreeConvertor)
	lon2 = (lon2 * degreeConvertor)
	s.location = Location{
		Latitude:  lat2,
		Longitude: lon2,
	}
	return s.location
}
