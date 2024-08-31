package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	Client   *mongo.Client
	Database *mongo.Database
}

func NewMongoDB(uri, dbName string) (*MongoDB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("erro ao conectar ao MongoDB: %v", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("erro ao pingar o MongoDB: %v", err)
	}

	fmt.Println("Conectado ao MongoDB!")

	return &MongoDB{
		Client:   client,
		Database: client.Database(dbName),
	}, nil
}

func (mdb *MongoDB) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := mdb.Client.Disconnect(ctx)
	if err != nil {
		return fmt.Errorf("erro ao desconectar do MongoDB: %v", err)
	}

	fmt.Println("Conexão com MongoDB fechada.")
	return nil
}

func (mdb *MongoDB) GetCollection(collectionName string) *mongo.Collection {
	return mdb.Database.Collection(collectionName)
}

func exampleUse() {
	mongoDB, err := NewMongoDB("mongodb://localhost:27017", "IncomingRawData")
	if err != nil {
		log.Fatal(err)
	}
	defer mongoDB.Close()

	// Acessa uma coleção
	collection := mongoDB.GetCollection("IncomingRawItems")

	// Exemplo: busca um documento aleatório na coleção
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var result map[string]interface{}
	err = collection.FindOne(ctx, map[string]interface{}{}).Decode(&result)
	if err != nil {
		log.Fatal("Erro ao buscar documento:", err)
	}

	fmt.Println("Documento encontrado:", result)
}
