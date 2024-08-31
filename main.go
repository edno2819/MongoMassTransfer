package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/edno2819/go-examples/src/database"
	middlewares "github.com/edno2819/go-examples/src/handlers"
	"github.com/edno2819/go-examples/src/utils"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func setup() {
	middlewares.DotEnvVariable("")
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	utils.PrintBeaultifull()
}

func getStringFromID(id interface{}) string {
	objectID, ok := id.(primitive.ObjectID)
	if !ok {
		fmt.Println("Error: _id is not of type ObjectID")
		return ""
	}
	return objectID.Hex()
}

type IncomingRawItem struct {
	ID   primitive.ObjectID `bson:"_id"`
	Name string             `bson:"name"`
	Age  int                `bson:"age"`
}

func convertToJsonString(result *mongo.SingleResult) (string, error) {
	var item bson.M
	err := result.Decode(&item)
	if err != nil {
		log.Println("Resultado do banco não encontrado!")
	}

	jsonData, err := json.Marshal(item)
	if err != nil {
		log.Println("Error converting to JSON")
		return "", err
	}
	return string(jsonData), nil
}

func convertToStruc(result *mongo.SingleResult, targetStruct interface{}) (interface{}, error) {
	err := result.Decode(targetStruct)
	if err != nil {
		log.Println("Error in convertion!")
		return nil, err
	}
	return targetStruct, nil
}

func testDb() {
	var result *mongo.SingleResult

	client := database.DbConnect()
	collection := client.Database("IncomingRawData").Collection("IncomingRawItems")

	result = collection.FindOne(context.TODO(), bson.M{})
	fmt.Println(convertToJsonString(result))

	result = collection.FindOne(context.TODO(), bson.M{})
	var item IncomingRawItem
	convertToStruc(result, &item)
}

func main() {
	setup()
	testDb()
	pathFile := "./stuck.csv"
	utils.ReadCSVInChunks(pathFile, 1000)

}
