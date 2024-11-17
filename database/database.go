package database

import (
	"context"
	"log"
	"os"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database
func ConnectDB() {
	uri := os.Getenv("MONGO_URI")
	ctx,cancel := context.WithTimeout(context.Background(),10*time.Second)
	defer cancel()

	client,err := mongo.Connect(ctx,options.Client().ApplyURI(uri))
	if err!=nil{
		log.Fatal(err)
	}

	if err:=client.Ping(ctx,nil) ;err!=nil{
		log.Fatal("MongoDB connection failed")
	}

	DB = client.Database(os.Getenv("DB_NAME"))
	log.Println("Connected to MongoDB")
}