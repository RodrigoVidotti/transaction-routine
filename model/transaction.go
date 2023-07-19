package model

type Transaction struct {
	TransactionId   int     `json:"transaction_id"`
	AccountId       int     `json:"account_id"`
	OperationTypeId int     `json:"operation_type_id"`
	Amount          float64 `json:"amount"`
	EventDatetime   string  `json:"event_datetime"`
}
