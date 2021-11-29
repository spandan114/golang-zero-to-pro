package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/spandan114/go-mongo-crud/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

//INSERT
func InsertOneData(movie model.Netflix) {
	result, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted value is :", result.InsertedID)
}

//UPDATE
func updateOneMovie(index string) {
	id, _ := primitive.ObjectIDFromHex(index)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("modified count: ", result.ModifiedCount)
}

//DELETE one movie
func deleteeOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}

	deleteCount, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MOvie got delete with delete count: ", deleteCount)
}

//DELETE all movie
func deleteAllMovies() int64 {

	deleteResult, err := collection.DeleteMany(context.Background(), bson.M{}, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("NUmber of movies delete: ", deleteResult.DeletedCount)
	return deleteResult.DeletedCount
}

//GET all movies
func getAllMovies() []primitive.M {
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())

	var movies []primitive.M
	for cursor.Next(context.Background()) {
		var movie bson.M
		err := cursor.Decode(movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}
	return movies
}

//CONTROLLER

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

}
