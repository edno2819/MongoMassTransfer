package process

import (
	"context"
	"log"
	"strconv"
	"sync"

	"github.com/edno2819/mongo-mass-transfer/src/database"
	"github.com/edno2819/mongo-mass-transfer/src/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Collection *mongo.Collection
	ChunkSize  int = 3000
)

func init() {
	ChunkSizeStr := utils.GetEnvVariableDef("CHUNK_SIZE", "50")
	value, err := strconv.Atoi(ChunkSizeStr)
	if err != nil {
		log.Printf("Error converting CHUNK_SIZE to int: %v", err)
		ChunkSize = 10
	}
	ChunkSize = value

	client := database.DbConnect()
	Collection = client.Database("Logs").Collection("TestGoMass")
}

func SaveChunk(chunk []interface{}, wg *sync.WaitGroup) {
	defer wg.Done()

	_, err := Collection.InsertMany(context.TODO(), chunk, options.InsertMany())
	if err != nil {
		log.Printf("Error inserting chunck %v", err)
		return
	}

	log.Printf("Salvando chunk de %d items", len(chunk))
}

func InsertChunk(dataProcessedChannel <-chan database.StockModel, wg *sync.WaitGroup) {
	var chunk []interface{}
	defer wg.Done()

	for data := range dataProcessedChannel {
		chunk = append(chunk, data)
		if len(chunk) >= ChunkSize {
			wg.Add(1)
			go SaveChunk(chunk, wg)
			chunk = make([]interface{}, 0)
		}
	}

	log.Println("Last chunck")

	// Whether exist the last chunck save its
	if len(chunk) > 0 {
		wg.Add(1)
		go SaveChunk(chunk, wg)
	}
}
