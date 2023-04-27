package controller

import (
	"go.mongodb.org/mongo-driver/mongo",
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
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("Mongo db connection success")
	collection = client.Database(dbName).Collection(colName)

	// 
	fmt.Println("Collection instance is ready")
}
