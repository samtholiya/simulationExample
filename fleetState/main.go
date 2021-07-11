package main

import (
	"log"
	"net/http"

	"github.com/fatih/color"
	"github.com/rs/cors"
	middlewares "github.com/samtholiya/fleetState/handlers"
	"github.com/samtholiya/fleetState/routes"
)

func main() {
	port := middlewares.DotEnvVariable("PORT")
	color.Cyan("🌏 Server running on localhost:" + port)

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	router := routes.Routes()
	c := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type", "Origin", "Accept", "*"},
	})

	handler := c.Handler(router)
	http.ListenAndServe(":"+port, middlewares.LogRequest(handler))
}
