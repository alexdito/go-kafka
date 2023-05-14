package model

type Transaction struct {
	User      string  `form:"user" binding:"required"`
	Amount    float64 `form:"amount" binding:"required"`
	Currency  string  `form:"currency" binding:"required"`
	Operation string  `from:"operation" binding:"required"`
	Timestamp int64
}

type TransactionList []Transaction
