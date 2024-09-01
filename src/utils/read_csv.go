package utils

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"sync"
)

var (
	Repeater int
)

func init() {
	repeaterStr := GetEnvVariableDef("REPETER", "50")
	value, err := strconv.Atoi(repeaterStr)
	if err != nil {
		log.Printf("Error converting REPETER to int: %v", err)
		Repeater = 10
	}
	Repeater = value
}

func ReadCSVChunks(filePath string, chunkSize int, channel chan interface{}) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("could not open file: %v", err)
	}
	defer file.Close()

	// Create a CSV reader
	reader := csv.NewReader(file)

	// Initialize a counter for the number of rows processed
	rowCount := 0

	for {
		// Read the next chunk of rows
		chunk := make([][]string, 0, chunkSize)
		for i := 0; i < chunkSize; i++ {
			record, err := reader.Read()
			if err != nil {
				if err.Error() == "EOF" {
					break
				}
				return fmt.Errorf("error reading row: %v", err)
			}
			chunk = append(chunk, record)
			rowCount++
		}

		if len(chunk) == 0 {
			break
		}

		// Process the chunk
		fmt.Printf("Processing chunk with %d rows\n", len(chunk))
		for _, row := range chunk {
			fmt.Println(row) // Example processing: print each row
		}

		// If we've reached the end of the file, break the loop
		if len(chunk) < chunkSize {
			break
		}
	}

	log.Printf("Total rows processed: %d\n", rowCount)
	return nil
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
