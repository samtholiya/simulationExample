package main

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/samtholiya/dataSimulator/simulator/vehicle"
)

func main() {
	numberOfCars := flag.Int("number", 21, "Number of simulators")
	hostURL := flag.String("server", "http://localhost:8080", "Server url")
	flag.Parse()
	envHostURL := os.Getenv("SERVER_URL")
	if envHostURL != "" {
		*hostURL = envHostURL
	}
	ctx, cancelFunc := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}
	wg.Add(*numberOfCars)
	for i := 0; i < *numberOfCars; i++ {
		simulator := vehicle.NewSimulator(fmt.Sprintf("simulated_%v", i), 52.20472+float64(i), 0.14056, 15+float64(i))
		go CarRunner(ctx, wg, *hostURL, simulator)
	}
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		cancelFunc()
		os.Exit(0)
	}()
	wg.Wait()
}

func CarRunner(ctx context.Context, wg *sync.WaitGroup, hostURL string, vehicle vehicle.Vehicle) {
	defer wg.Done()
	client := http.Client{
		Timeout: time.Second * 30,
	}
	finalURL := hostURL + "/vehicle/" + vehicle.VIN
	ticker := time.NewTicker(time.Minute)
	log.Printf("%v vehicle started", vehicle.VIN)
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
		}
		data, _ := json.Marshal(vehicle.Location())
		req, err := http.NewRequestWithContext(ctx, http.MethodPost, finalURL, bytes.NewReader(data))
		if err != nil {
			log.Fatalf("[ERROR] %v", err.Error())
		}
		res, err := client.Do(req)
		if err != nil {
			log.Fatalf("[ERROR] %v", err.Error())
			continue
		}
		data, _ = ioutil.ReadAll(res.Body)
		if res.StatusCode != http.StatusOK {
			log.Printf("%v: %v", vehicle.VIN, string(data))
		}
		res.Body.Close()
	}
}
