package db

import (
	"context"
	"log"

	"github.com/fatih/color"
	middlewares "github.com/samtholiya/fleetState/handlers"
	"github.com/samtholiya/fleetState/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Dbconnect -> connects mongo
func Dbconnect() *mongo.Client {
	mongoURL := middlewares.DotEnvVariable("MONGO_URL")
	color.Green("[INFO] Mongo URL: %v", mongoURL)
	clientOptions := options.Client().ApplyURI(mongoURL)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	models.CreateModels(client)
	if err != nil {
		log.Fatal("⛒ Connection Failed to Database")
		log.Fatal(err)
	}
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("⛒ Connection Failed to Database")
		log.Fatal(err)
	}
	color.Green("⛁ Connected to Database")
	return client
}
