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

func init() {
	middlewares.DotEnvVariable("")
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	utils.PrintBeaultifull()
}

func ReadCSV(filePath string, chunkSize int, channel chan interface{}, wg *sync.WaitGroup) {
	defer wg.Done()

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("could not open file: %v", err)
	}
	defer file.Close()

	// Create a CSV reader
	reader := csv.NewReader(file)

	rowCount := 0
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

	log.Printf("Total rows Loadded: %d\n", rowCount)
	close(channel)
}

func processRow(rowsChannel chan interface{}, dataProcessedChannel chan interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for data := range rowsChannel {
		dataProcessedChannel <- data
	}
}

func saveChunk(chunk []interface{}, qtdSaved *int, wg *sync.WaitGroup) {
	defer wg.Done()

	time.Sleep(1200 * time.Millisecond)
	*qtdSaved += len(chunk)
	log.Printf("Salvando chunk de %d items", len(chunk))
}

func insertChunck(dataProcessedChannel chan interface{}, wg *sync.WaitGroup, qtdSaved *int) {
	const chunkSize = 500
	var chunk []interface{}
	defer wg.Done()

	for data := range dataProcessedChannel {
		chunk = append(chunk, data)
		if len(chunk) >= chunkSize {
			wg.Add(1)
			go saveChunk(chunk, qtdSaved, wg)
			chunk = make([]interface{}, 0)
		}
	}

	if len(chunk) > 0 {
		wg.Add(1)
		go saveChunk(chunk, qtdSaved, wg)
	}
}

func main() {
	const MaxSizeRowsBuffer int16 = 500
	const MaxSizeDataProcessedBuffer int16 = 50
	var qtdSaved int = 0

	rowsChannel := make(chan interface{}, MaxSizeRowsBuffer)
	dataProcessedChannel := make(chan interface{}, MaxSizeDataProcessedBuffer)

	pathFile := "tmp/stock.csv"
	var wg sync.WaitGroup

	wg.Add(1)
	go ReadCSV(pathFile, 1000, rowsChannel, &wg)

	for c := 0; c < runtime.NumCPU()-2; c++ {
		wg.Add(1)
		go processRow(rowsChannel, dataProcessedChannel, &wg)
	}

	wg.Add(1)
	go insertChunck(rowsChannel, &wg, &qtdSaved)

	wg.Wait()
	defer close(dataProcessedChannel)
	log.Println("Data Saved", qtdSaved)
}
