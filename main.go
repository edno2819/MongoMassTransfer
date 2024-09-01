package main

import (
	"log"
	"runtime"
	"sync"

	"github.com/edno2819/mongo-mass-transfer/src/database"
	"github.com/edno2819/mongo-mass-transfer/src/process"
	"github.com/edno2819/mongo-mass-transfer/src/utils"
	"github.com/joho/godotenv"
)

const (
	MaxSizeRowsBuffer          int16  = 5000
	MaxSizeDataProcessedBuffer int16  = 5000
	PathFile                   string = "tmp/stock.csv"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	utils.PrintBeaultifull()
}

func main() {
	rowsChannel := make(chan interface{}, MaxSizeRowsBuffer)
	dataProcessedChannel := make(chan database.StockModel, MaxSizeDataProcessedBuffer)

	var wg sync.WaitGroup
	var wgSave sync.WaitGroup

	wg.Add(1)
	go utils.ReadCSV(PathFile, rowsChannel, &wg)

	numWorkers := runtime.NumCPU()
	for c := 0; c < numWorkers; c++ {
		wg.Add(1)
		go process.ProcessRow(rowsChannel, dataProcessedChannel, &wg)
	}

	wgSave.Add(1)
	go process.InsertChunk(dataProcessedChannel, &wgSave)

	wg.Wait()
	close(dataProcessedChannel)

	wgSave.Wait()
}
