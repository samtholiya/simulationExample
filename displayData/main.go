package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Response struct {
	Data struct {
		Location struct {
			Latitude  float64
			Longitude float64
		}
		Speed float64
		VIN   string
	}
	StatusCode int
	Message    string
}

func main() {
	vin := flag.String("vin", "", "Vin of vehicle")
	server := flag.String("server", "http://localhost:8080", "Server endpoint")
	flag.Parse()
	for ; ; time.Sleep(time.Minute) {
		res, err := http.Get(fmt.Sprintf("%v/vehicle/%v/stream", *server, *vin))
		if err != nil {
			log.Println(err)
			continue
		}
		var result Response
		if res.StatusCode != http.StatusOK {
			log.Printf("Status code %v", res.StatusCode)
			continue
		}
		err = json.NewDecoder(res.Body).Decode(&result)
		if err != nil {
			log.Println(err)
			continue
		}
		res.Body.Close()
		fmt.Printf("vin: %v is currently moving with a speed of %.2f m/s and is at location: %v. %v\n", result.Data.VIN, result.Data.Speed, result.Data.Location.Latitude, result.Data.Location.Longitude)

	}
}
