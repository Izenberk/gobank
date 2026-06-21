package domain

import "time"


type Account struct {
	ID				int64
	Fullname	string
	Balance		int64
	CreatedAt time.Time
}

type Transaction struct {
	ID				int64
	AccountID	int64
	Type			TransactionType
	Amount		int64
	Timestamp time.Time
}

type TransactionType string

const (
	Deposit			TransactionType = "deposit"
	Withdraw		TransactionType = "withdraw"
)
