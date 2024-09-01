package transformer

import "github.com/edno2819/mongo-mass-transfer/src/database"

func FormaterDataStock(row []string) database.StockModel {
	return database.StockModel{
		PartnerId:             row[0],
		ItemPartnerInstoreSKU: row[1],
		ItemPartnerInstoreId:  row[2],
		ProcessId:             0,
		OrderPartnerData:      nil,
	}
}
