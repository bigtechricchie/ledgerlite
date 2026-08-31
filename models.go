package main

type TransactionType string

const (
	Deposit    TransactionType = "DEPOSIT"
	Withdrawal TransactionType = "WITHDRAWAL"
)

type Account struct {
	ID   string
	Name string
}

type Transaction struct {
	ID          string
	AccountID   string
	Type        TransactionType
	Amount      int64 // Amount is stored in minor units (pence).
	Description string
}
