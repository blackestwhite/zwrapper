package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func Connect() {
	log.Println("Database connecting...")

	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://mongodb:27017")

	// Connect to MongoDB
	connCtx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	client, err := mongo.Connect(connCtx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	Client = client

	// Check the connection
	pingCtx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	err = Client.Ping(pingCtx, nil)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database Connected.")
}

func Disconnect() {
	if Client != nil {
		if err := Client.Disconnect(context.Background()); err != nil {
			log.Fatal("Error closing database connection:", err)
		} else {
			log.Println("Database connection closed.")
		}
	}
}
