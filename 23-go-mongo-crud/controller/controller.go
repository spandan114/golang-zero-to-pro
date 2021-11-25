package controller

import (
	"context"
	"fmt"
	"log"

	"github.com/spandan114/go-mongo-crud/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const mongoURI string = "mongodb+srv://spandan:spandan1234@cluster0.3cnzz.mongodb.net/go-crud?retryWrites=true&w=majority"
const dbName string = "go-crud"
const collectionName string = "netflix"

var collection *mongo.Collection

func ConnectMongodb() {
	//client option
	clintOption := options.Client().ApplyURI(mongoURI)

	//connect with mongodb
	client, err := mongo.Connect(context.TODO(), clintOption)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Mongodb connected successfully")

	collection = client.Database(dbName).Collection(collectionName)

	//collection instance
	fmt.Println("Collection instance is ready")
}

//MONGODB helpers

func InsertOneData(movie model.Netflix) {
	result, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted value is :", result.InsertedID)
}
