package driver

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"time"
)

func GetClient() (client *mongo.Client) {

	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		panic("MONGO_URI is not set or empty")
	}

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(os.Getenv("MONGO_URI")).SetServerAPIOptions(serverAPI)

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	client, err := mongo.Connect(ctx, opts)

	if err != nil {
		cancel()
		panic(err)
	}

	var result bson.M

	if err := client.Database("search").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		panic(err)
	}

	return

}

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {

	collection := client.Database("SimpleGroqImp").Collection(collectionName)

	return collection
}

func CloseClient(client *mongo.Client) {

	ctx := context.TODO()
	err := client.Disconnect(ctx)
	if err != nil {
		panic(err)
	}
}
