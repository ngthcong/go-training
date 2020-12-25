package repositories

import (
	"context"
	"fmt"
	"log"


	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

)

var uri  = "mongodb://localhost:27017"
var ctx = context.TODO()
var dbName = "db_school"

func NewMongoDB() *mongo.Database{
	// Replace the uri string with your MongoDB deployment's connection string.
	clientOptions := options.Client().ApplyURI(uri)
	var client, err = mongo.Connect(ctx,clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Ping the primary
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal(err)
	}
	db := client.Database(dbName)
	fmt.Println("Successfully connected and pinged.")
	return db
}
