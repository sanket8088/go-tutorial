package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/sanket/mongoapi/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://snihal:y1gsmzvCP4YvZbii@demlolocal.8inp84j.mongodb.net/?retryWrites=true&w=majority"
const dbName = "netflix"
const colName = "watchlist"

var collection *mongo.Collection

// connect with mongoDB

func init() {
	// client option
	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Mongo db connection success")
	collection = client.Database(dbName).Collection(colName)

	//
	fmt.Println("Collection instance is ready")
}

// MONGODB helpers

// insert 1 movie
func insertOneMovie(movie model.Netflix) {
	m := movie
	m.Watched = false
	fmt.Println(movie)
	inserted, err := collection.InsertOne(context.Background(), m)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(inserted)
	fmt.Println("Inserted one Movie in db with id", inserted.InsertedID)
}

// update 1 record

func updateOneMovie(movieId string) {
	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
	fmt.Println("modified count", result.ModifiedCount)

}

// delete 1 record

func deleteOneMovie(movieId string) {
	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": id}
	result, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Movie got deleted with delete count:", result)

}

// delete all record

func deleteAllMovies() int64 {
	deleteCount, err := collection.DeleteMany(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(deleteCount)
	fmt.Println("Movies got deleted", deleteCount.DeletedCount)
	return deleteCount.DeletedCount

}

func getAllMovies() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	var movies []primitive.M
	for cur.Next(context.Background()) {
		var movie bson.M
		err := cur.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}
	defer cur.Close(context.Background())
	return movies
}

// Actual controller

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var singleMovie model.Netflix
	if err := json.NewDecoder(r.Body).Decode(&singleMovie); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	validate := validator.New()
	err := validate.Struct(singleMovie)
	if err != nil {
		fmt.Println(err)
		fmt.Println(err.Error())
		http.Error(w, "Keys missing", http.StatusBadRequest)
		return
	}
	_ = json.NewDecoder(r.Body).Decode(&singleMovie)
	insertOneMovie(singleMovie)
	json.NewEncoder(w).Encode(singleMovie)
}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")
	params := mux.Vars(r)
	var singleMovie model.Netflix
	if err := json.NewDecoder(r.Body).Decode(&singleMovie); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])

}

func DeleteOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	params := mux.Vars(r)
	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])

}

func DeleteAllMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	count := deleteAllMovies()
	json.NewEncoder(w).Encode(count)

}
