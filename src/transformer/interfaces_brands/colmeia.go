package transformer

import (
	"strconv"
	"strings"
)

type FormatStockColmeia struct {
	ID_ECOMERCE int
}

func NewFormatStockColmeia() *FormatStockColmeia {
	return &FormatStockColmeia{
		ID_ECOMERCE: 15,
	}
}

func (f *FormatStockColmeia) GetOrderID(data map[string]interface{}) string {
	if val, ok := data["CD_PRODUTO"]; ok {
		return strconv.Itoa(val.(int)) // supondo que CD_PRODUTO seja um int
	}
	return ""
}

func (f *FormatStockColmeia) GetStoreID(data map[string]interface{}) *int {
	if val, ok := data["LOCALIDADE"]; ok {
		value := val.(string)
		parts := strings.Split(value, "-")
		if len(parts) > 0 {
			id, err := strconv.Atoi(strings.TrimSpace(parts[0]))
			if err == nil {
				return &id
			}
		}
	}
	return nil
}
