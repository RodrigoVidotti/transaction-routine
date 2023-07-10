package model

type Transaction struct {
	TransactionId   string `json:"transaction_id"`
	AccountId       string `json:"account_id"`
	OperationTypeId string `json:"operation_type_id"`
	Amount          string `json:"amount"`
	EventDatetime   string `json:"event_datetime"`
}
