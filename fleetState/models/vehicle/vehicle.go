package vehicle

import (
	"context"
	"time"

	"github.com/samtholiya/fleetState/models/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	collectionName = "vehicle"
)

type Vehicle struct {
	ID               primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	PreviousLocation Location           `bson:"previousLocation"`
	CurrentLocation  Location           `bson:"currentLocation"`
	VIN              string             `bson:"vin"`
	CurrentUpdate    time.Time          `bson:"currentUpdate"`
	PreviousUpdate   time.Time          `bson:"previousUpdate"`
}

func FindOne(ctx context.Context,
	client *mongo.Client,
	filter interface{},
	opts ...*options.FindOneOptions) (Vehicle, bool) {
	var vehicle Vehicle
	collection := client.Database(config.DatabaseName).Collection(collectionName)
	result := collection.FindOne(ctx, filter, opts...)
	if result.Err() != nil {
		return Vehicle{}, false
	}
	result.Decode(&vehicle)
	return vehicle, true
}

func InsertOne(ctx context.Context, client *mongo.Client, vehicle Vehicle) (*mongo.InsertOneResult, error) {

	collection := client.Database(config.DatabaseName).Collection(collectionName)
	date := time.Now()
	vehicle.CurrentUpdate = date
	vehicle.PreviousUpdate = vehicle.CurrentUpdate
	return collection.InsertOne(ctx, vehicle)
}

func UpdateOne(ctx context.Context, client *mongo.Client, vehicle Vehicle) (*mongo.UpdateResult, error) {
	collection := client.Database(config.DatabaseName).Collection(collectionName)
	return collection.UpdateOne(ctx, bson.M{"_id": vehicle.ID},
		bson.D{
			primitive.E{Key: "$set", Value: bson.D{
				primitive.E{Key: "currentLocation", Value: vehicle.CurrentLocation},
				primitive.E{Key: "previousLocation", Value: vehicle.PreviousLocation},
				primitive.E{Key: "previousUpdate", Value: vehicle.CurrentUpdate},
				primitive.E{Key: "currentUpdate", Value: time.Now().Unix()},
			}},
		})
}

func CreateCollection(ctx context.Context, client *mongo.Client) {
	collection := client.Database(config.DatabaseName).Collection(collectionName)
	collection.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.M{
			"vin": 1,
		},
	})
}
