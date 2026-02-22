package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sidz111/go-gin-emp-crud/controller"
	"github.com/sidz111/go-gin-emp-crud/db"
	"github.com/sidz111/go-gin-emp-crud/repository"
	"github.com/sidz111/go-gin-emp-crud/service"
)

func main() {
	db := db.ConnectDb()

	defer db.Close()

	router := gin.Default()

	repo := repository.NewEmployeeRepository(db)
	serv := service.NewEmployeeService(repo)
	controller := controller.NewEmployeeController(serv)

	r := router.Group("/employees")
	{
		r.POST("/", controller.CreateEmployee)
		r.GET("/:id", controller.GetEmployeeById)
		r.GET("/", controller.GetAllEmployees)
		r.PUT("/", controller.UpdateEmployee)
		r.DELETE("/:id", controller.DeleteEmployee)
	}
	router.Run(":8080")

}
