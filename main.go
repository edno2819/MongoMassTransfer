package main

import (
	"context"
	"log"
	"runtime"
	"sync"

	"github.com/edno2819/go-examples/src/database"
	"github.com/edno2819/go-examples/src/utils"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	Repeater                   int    = 50
	ChunkSize                  int    = 3000
	MaxSizeRowsBuffer          int16  = 5000
	MaxSizeDataProcessedBuffer int16  = 5000
	PathFile                   string = "tmp/stock.csv"
)

var (
	qtdSaved   int
	mu         sync.Mutex
	Collection *mongo.Collection
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	utils.PrintBeaultifull()
}

func getCollection() *mongo.Collection {
	client := database.DbConnect()
	return client.Database("Logs").Collection("TestGoMass")
}

func processRow(rowsChannel <-chan interface{}, dataProcessedChannel chan<- interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for data := range rowsChannel {
		row := data.([]string)
		// Mapeia os campos para os nomes das colunas do MongoDB (ajuste conforme necessário)
		document := map[string]interface{}{
			"data":      row[0],
			"loja":      row[1],
			"idProduct": row[2],
		}
		dataProcessedChannel <- document
	}
}

func saveChunk(chunk []interface{}, wg *sync.WaitGroup) {
	defer wg.Done()

	_, err := Collection.InsertMany(context.TODO(), chunk, options.InsertMany())
	if err != nil {
		log.Printf("Error inserting chunck %v", err)
		return
	}

	mu.Lock()
	qtdSaved += len(chunk)
	mu.Unlock()
	log.Printf("Salvando chunk de %d items", len(chunk))
}

func insertChunk(dataProcessedChannel <-chan interface{}, wg *sync.WaitGroup) {
	var chunk []interface{}
	defer wg.Done()

	for data := range dataProcessedChannel {
		chunk = append(chunk, data)
		if len(chunk) >= ChunkSize {
			wg.Add(1)
			go saveChunk(chunk, wg)
			chunk = make([]interface{}, 0)
		}
	}

	log.Println("Last chunck")

	// Salva o último chunk, se houver
	if len(chunk) > 0 {
		wg.Add(1)
		go saveChunk(chunk, wg)
	}
}

func main() {
	Collection = getCollection()

	rowsChannel := make(chan interface{}, MaxSizeRowsBuffer)
	dataProcessedChannel := make(chan interface{}, MaxSizeDataProcessedBuffer)

	var wg sync.WaitGroup
	var wgSave sync.WaitGroup

	wg.Add(1)
	go utils.ReadCSV(PathFile, rowsChannel, &wg)

	numWorkers := runtime.NumCPU()
	for c := 0; c < numWorkers; c++ {
		wg.Add(1)
		go processRow(rowsChannel, dataProcessedChannel, &wg)
	}

	wgSave.Add(1)
	go insertChunk(dataProcessedChannel, &wgSave)

	wg.Wait()
	close(dataProcessedChannel)

	wgSave.Wait()
	log.Println("Data Saved", qtdSaved)
}
