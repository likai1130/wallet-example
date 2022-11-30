package v1

type Transaction struct {
}

type TransactionList struct {
	Items []*Transaction `json:"items"`
}
