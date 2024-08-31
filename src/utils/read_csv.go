package utils

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func ReadCSVInChunks(filePath string, chunkSize int) error {
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
