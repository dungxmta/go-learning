package driver

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

type MongoDB struct {
	Client *mongo.Client
}

var Mongo = &MongoDB{}

func ConnectMongoDB(uri string) *MongoDB {
	// new client
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	// set timeout when connect
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}

	// ping test
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}

	fmt.Println("connect db ok!")

	Mongo.Client = client
	return Mongo
}
