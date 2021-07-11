package vehicle

import "testing"

func TestCarLocation(t *testing.T) {
	vehicleObj := NewSimulator("MH04AC8899", 52.20472, 0.14056, 15)
	location := vehicleObj.Location()
	if location.Latitude != 52.20462299620793 || location.Longitude != 0.360433887489931 {
		t.Errorf("Latitude: %v Longitude: %v", location.Latitude, location.Longitude)
	}
	location = vehicleObj.Location()
	if location.Latitude != 52.20452599313012 || location.Longitude != 0.5803072949950739 {
		t.Errorf("Latitude: %v Longitude: %v", location.Latitude, location.Longitude)
	}
}
