package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func DBInstance() *mongo.Client {

	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Warning: unable to .env file")
	}
	MongoDb := os.Getenv("MONGO_URI")
	if MongoDb == "" {
		log.Fatal("MONGO_URI environment variable not set", MongoDb)
	}
	clientOptions := options.Client().ApplyURI(MongoDb)
	client, err := mongo.Connect(clientOptions)

	if err != nil {
		return nil
	}
	return client
}

var Client *mongo.Client = DBInstance()

func OpenCollection(collectionName string) *mongo.Collection {

	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Warning: unable to find .env file")
	}
	databaseName := os.Getenv(("DATABASE_NAME"))
	fmt.Println("Database Name:", databaseName)

	collection := Client.Database(databaseName).Collection(collectionName)

	if collection == nil {
		return nil
	}
	return collection

}
