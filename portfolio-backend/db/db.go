package db

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database
var client *mongo.Client

func ConnectDB() {
	godotenv.Load()

	mongoURI := os.Getenv("MONGO_URI")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var err error

	client, err = mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))

	if err != nil {
		log.Fatal(err)
	}

	DB = client.Database(os.Getenv("DB_NAME"))

	log.Println("MongoDB Atlas Cluster Connected")
}

func GetClient() *mongo.Client {
	return  client
}