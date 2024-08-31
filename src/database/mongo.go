package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DbConnect() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	fmt.Println()

	clientOptions := options.Client().ApplyURI(os.Getenv("UPLAN_URI_MONGO"))

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("Falha ao conectar ao MongoDB:", err)
	}

	log.Println("Conectado ao MongoDB!")
	return client

}

func Examples(client *mongo.Client) {
	db := client.Database("IncomingRawData")
	collection := db.Collection("IncomingRawItems")

	pipeline := mongo.Pipeline{
		bson.D{{"$sample", bson.D{{"size", 1}}}},
	}

	cur, err := collection.Aggregate(context.TODO(), pipeline)
	if err != nil {
		log.Fatal("Erro ao buscar documento aleatório:", err)
	}
	defer cur.Close(context.TODO())

	if cur.Next(context.TODO()) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal("Erro ao decodificar documento:", err)
		}
		fmt.Println("Documento aleatório encontrado:", result)
	} else {
		fmt.Println("Nenhum documento encontrado.")
	}
	// Fecha a conexão com o MongoDB
	err = client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal("Erro ao desconectar do MongoDB:", err)
	}

	fmt.Println("Conexão com MongoDB fechada.")
}
