package model

type Employee struct {
	Name    string  `json:"name" binding:"required"`
	Address string  `json:"address" binding:"required"`
	Salary  float64 `json:"salary" binding:"required"`
}
