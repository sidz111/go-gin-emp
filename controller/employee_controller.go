package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sidz111/go-gin-emp-crud/model"
	"github.com/sidz111/go-gin-emp-crud/service"
)

type EmployeeController struct {
	service *service.EmployeeService
}

func NewEmployeeController(service *service.EmployeeService) *EmployeeController {
	return &EmployeeController{service: service}
}

func (c *EmployeeController) CreateEmployee(ctx *gin.Context) {
	var emp model.Employee

	if err := ctx.ShouldBindJSON(&emp); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := c.service.CreateEmployee(ctx.Request.Context(), &emp); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"message":  "employee created successfully",
		"employee": emp,
	})
}

func (c *EmployeeController) GetEmployeeById(ctx *gin.Context) {
	id := ctx.Param("id")
	newId, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}
	emp, err := c.service.GetEmployeeById(ctx.Request.Context(), newId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"employee": emp,
	})
}

func (c *EmployeeController) GetAllEmployees(ctx *gin.Context) {
	employees, err := c.service.GetAllEmployees(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"employees": employees,
	})
}

func (c *EmployeeController) UpdateEmployee(ctx *gin.Context) {
	var emp model.Employee
	if err := ctx.ShouldBindJSON(&emp); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}

	if err := c.service.UpdateEmployee(ctx.Request.Context(), &emp); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message":  "employee successfully updated",
		"employee": emp,
	})
}

func (c *EmployeeController) DeleteEmployee(ctx *gin.Context) {
	id := ctx.Param("id")
	newId, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}
	if err := c.service.DeleteEmployeeById(ctx.Request.Context(), newId); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "employee deleted successfully",
	})
}
