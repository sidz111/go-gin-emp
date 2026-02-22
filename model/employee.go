package model

type Employee struct {
	ID      int     `json:"id"`
	Name    string  `json:"name" binding:"required"`
	Address string  `json:"address" binding:"required"`
	Salary  float64 `json:"salary" binding:"required"`
}
