package models

import (
	"context"

	"github.com/samtholiya/fleetState/models/vehicle"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateModels(client *mongo.Client) {
	ctx := context.TODO()
	vehicle.CreateCollection(ctx, client)
}
