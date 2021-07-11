package routes

import (
	"github.com/gorilla/mux"
	"github.com/samtholiya/fleetState/controllers"
)

// Routes -> define endpoints
func Routes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/vehicle/{vin}", controllers.CreateUpdateVehicleEndpoint).Methods("POST")
	router.HandleFunc("/vehicle/{vin}/stream", controllers.GetVehicleEndpoint).Methods("GET")
	return router
}
