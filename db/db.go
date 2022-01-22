package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

func formConnection() (*mongo.Client, mongo.context) {
	// Set context and connection string -- Set <password>
	clientOptions := mongo.options.Client().
		ApplyURI("mongodb+srv://ReamerDB:<password>@reamerdb.jgbyr.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to AWS mongodb
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	return client, ctx
}

func dbConnection() *mongo.Client {

	// Connect to DB
	client, context := formConnection()

	// Verify connection
	err := &client.Ping(context, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to AWS MongoDB\n")

	return client
}
