package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sidz111/go-gin-emp-crud/controller"
	"github.com/sidz111/go-gin-emp-crud/db"
	"github.com/sidz111/go-gin-emp-crud/repository"
	"github.com/sidz111/go-gin-emp-crud/service"
)

func main() {
	database := db.ConnectDb()

	defer database.Close()

	router := gin.Default()

	repo := repository.NewEmployeeRepository(database)
	serv := service.NewEmployeeService(repo)
	controller := controller.NewEmployeeController(serv)

	api := router.Group("/employees")
	{
		api.POST("/", controller.CreateEmployee)
		api.GET("/:id", controller.GetEmployeeById)
		api.GET("/", controller.GetAllEmployees)
		api.PUT("/", controller.UpdateEmployee)
		api.DELETE("/:id", controller.DeleteEmployee)
	}
	router.Run(":8080")

}
