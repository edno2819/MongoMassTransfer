package process

import (
	"sync"

	"github.com/edno2819/mongo-mass-transfer/src/database"
	"github.com/edno2819/mongo-mass-transfer/src/transformer"
)

func ProcessRow(rowsChannel <-chan interface{}, dataProcessedChannel chan<- database.StockModel, wg *sync.WaitGroup) {
	defer wg.Done()
	for data := range rowsChannel {
		row := data.([]string)
		dataProcessedChannel <- transformer.FormaterDataStock(row)
	}
}
