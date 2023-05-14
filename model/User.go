package model

type User struct {
	User   string `form:"user" binding:"required"`
	Amount int
}
