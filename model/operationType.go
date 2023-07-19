package model

type OperationType struct {
	OperationTypeID int    `json:"operation_type_id"`
	Name            string `json:"name"`
}

const (
	Debit  = -1
	Credit = 1
)

var ValidOperationTypes = map[int]int{
	1: Debit,  // COMPRA A VISTA
	2: Debit,  // COMPRA PARCELADA
	3: Debit,  // SAQUE
	4: Credit, // PAGAMENTO
}
