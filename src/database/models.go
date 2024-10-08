package database

type StockModel struct {
	PartnerId             string      `json:"partnerId,omitempty" bson:"partnerId,omitempty" validate:"required,alpha"`
	ItemPartnerInstoreSKU string      `json:"itemPartnerInstoreSKU,omitempty" bson:"itemPartnerInstoreSKU,omitempty" validate:"required,alpha"`
	ItemPartnerInstoreId  string      `json:"itemPartnerInstoreId,omitempty" bson:"itemPartnerInstoreId,omitempty" validate:"required,alpha"`
	ProcessId             int32       `json:"processId,omitempty" bson:"processId,omitempty" validate:"required,alpha"`
	OrderPartnerData      interface{} `json:"orderPartnerData,omitempty" bson:"orderPartnerData,omitempty" validate:"required,alpha"`
}
