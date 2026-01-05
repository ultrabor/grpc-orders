package model

type Order struct {
	ID     int64
	UserID string
	Amount float64
	Status string
}
