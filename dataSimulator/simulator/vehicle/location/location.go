package location

type Location struct {
	Latitude  float64
	Longitude float64
}

func (l Location) GetLatitude() float64 {
	return l.Latitude
}

func (l Location) GetLongitude() float64 {
	return l.Longitude
}
