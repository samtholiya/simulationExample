package vehicle

import "math"

const (
	radiusOfEarth = 6378.1
)

type Location struct {
	Latitude  float64 `json:"Latitude" bson:"latitude" validate:"required"`
	Longitude float64 `json:"Longitude" bson:"longitude" validate:"required"`
}

// GetDistance returns the distance in km
func (l Location) GetDistance(otherLocation Location) float64 {
	lat1 := l.Latitude
	lon1 := l.Longitude
	lat2 := otherLocation.Latitude
	lon2 := otherLocation.Longitude

	dLat := deg2rad(lat2 - lat1) // deg2rad below
	dLon := deg2rad(lon2 - lon1)
	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(deg2rad(lat1))*math.Cos(deg2rad(lat2))*
			math.Sin(dLon/2)*math.Sin(dLon/2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	distance := radiusOfEarth * c // Distance in km
	return distance

}

// TODO: Should be somewhere else
func deg2rad(deg float64) float64 {
	return deg * (math.Pi / 180)
}
