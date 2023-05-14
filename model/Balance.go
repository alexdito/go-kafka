package model

type Balance struct {
	User   string  `form:"user" binding:"required"`
	Amount float64 `form:"amount" binding:"required"`
}
