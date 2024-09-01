package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"runtime"
	"sync"
	"time"

	middlewares "github.com/edno2819/go-examples/src/handlers"
	"github.com/edno2819/go-examples/src/utils"
	"github.com/joho/godotenv"
)

const Repeater int = 10
const ChunkSize int = 500
const MaxSizeRowsBuffer int16 = 500
const MaxSizeDataProcessedBuffer int16 = 50
const PathFile string = "tmp/stock.csv"

func init() {
	middlewares.DotEnvVariable("")
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	utils.PrintBeaultifull()
}

func ReadCSV(filePath string, channel chan<- interface{}, wg *sync.WaitGroup) {
	defer wg.Done()

	rowCount := 0
	for i := 0; i < Repeater; i++ {
		file, err := os.Open(filePath)
		if err != nil {
			log.Fatalf("could not open file: %v", err)
			return
		}

		reader := csv.NewReader(file)
		for {
			rec, err := reader.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal(err)
			}
			channel <- rec
			rowCount++
		}

		file.Close()
	}

	log.Printf("Total rows Loaded: %d\n", rowCount)
	close(channel)
}

func processRow(rowsChannel <-chan interface{}, dataProcessedChannel chan<- interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for data := range rowsChannel {
		time.Sleep(70 * time.Microsecond)
		dataProcessedChannel <- data
	}
}

func saveChunk(chunk []interface{}, qtdSaved *int, wg *sync.WaitGroup) {
	defer wg.Done()

	time.Sleep(1200 * time.Millisecond)
	*qtdSaved += len(chunk)
	log.Printf("Salvando chunk de %d items", len(chunk))
}

func insertChunk(dataProcessedChannel <-chan interface{}, wg *sync.WaitGroup, qtdSaved *int) {
	var chunk []interface{}
	defer wg.Done()

	for data := range dataProcessedChannel {
		chunk = append(chunk, data)
		if len(chunk) >= ChunkSize {
			wg.Add(1)
			go saveChunk(chunk, qtdSaved, wg)
			chunk = make([]interface{}, 0)
		}
	}

	log.Println("Last chunck")

	// Salva o último chunk, se houver
	if len(chunk) > 0 {
		wg.Add(1)
		go saveChunk(chunk, qtdSaved, wg)
	}
}

func main() {
	var qtdSaved int = 0

	rowsChannel := make(chan interface{}, MaxSizeRowsBuffer)
	dataProcessedChannel := make(chan interface{}, MaxSizeDataProcessedBuffer)

	var wg sync.WaitGroup
	var wgSave sync.WaitGroup

	wg.Add(1)
	go ReadCSV(PathFile, rowsChannel, &wg)

	numWorkers := runtime.NumCPU()
	for c := 0; c < numWorkers; c++ {
		wg.Add(1)
		go processRow(rowsChannel, dataProcessedChannel, &wg)
	}

	wgSave.Add(1)
	go insertChunk(dataProcessedChannel, &wgSave, &qtdSaved)

	wg.Wait()
	close(dataProcessedChannel)

	wgSave.Wait()
	log.Println("Data Saved", qtdSaved)
}
