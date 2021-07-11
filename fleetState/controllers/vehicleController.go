package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/samtholiya/fleetState/db"
	middlewares "github.com/samtholiya/fleetState/handlers"
	"github.com/samtholiya/fleetState/models/vehicle"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var client = db.Dbconnect()

// CreateUpdateVehicleEndpoint -> create vehicle
var CreateUpdateVehicleEndpoint = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	ctx := context.TODO()
	vin := params["vin"]
	var location vehicle.Location
	err := json.NewDecoder(request.Body).Decode(&location)
	if err != nil {
		middlewares.ServerErrResponse(err.Error(), response)
		return
	}
	vehicleObject, isExisting := vehicle.FindOne(ctx, client, bson.D{primitive.E{Key: "vin", Value: vin}})
	if !isExisting {
		result, err := vehicle.InsertOne(ctx, client, vehicle.Vehicle{
			VIN:              vin,
			PreviousLocation: location,
			CurrentLocation:  location,
		})
		if err != nil {
			middlewares.ServerErrResponse(err.Error(), response)
		}
		res, _ := json.Marshal(result.InsertedID)
		middlewares.SuccessResponse(`Inserted at `+strings.Replace(string(res), `"`, ``, 2), response)
		return
	}
	vehicleObject.PreviousLocation = vehicleObject.CurrentLocation
	vehicleObject.CurrentLocation = location

	res, err := vehicle.UpdateOne(ctx, client, vehicleObject)
	if err != nil {
		middlewares.ServerErrResponse(err.Error(), response)
		return
	}
	middlewares.SuccessResponse("Updated "+strconv.Itoa(int(res.ModifiedCount)), response)
})

var GetVehicleEndpoint = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	ctx := context.TODO()
	vin := params["vin"]

	vehicleObject, isExisting := vehicle.FindOne(ctx, client, bson.D{primitive.E{Key: "vin", Value: vin}})
	if !isExisting {
		middlewares.ServerErrResponse("No Data with given vin", response)
		return
	}
	distance := vehicleObject.CurrentLocation.GetDistance(vehicleObject.PreviousLocation)
	diffTime := vehicleObject.CurrentUpdate.Sub(vehicleObject.PreviousUpdate)
	seconds := diffTime.Seconds()
	speed := distance / seconds
	if seconds == 0 {
		speed = 0
	}
	middlewares.SuccessRespond(map[string]interface{}{
		"VIN":      vehicleObject.VIN,
		"Location": vehicleObject.CurrentLocation,
		"Speed":    speed,
	}, response)
})
