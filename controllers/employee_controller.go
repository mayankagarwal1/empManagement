package controllers

import (
    "empManagement/config"
    "empManagement/models"
    "github.com/gin-gonic/gin"
    "net/http"
)

func CreateEmployee(c *gin.Context) {
    // Bind the request JSON data to the Employee model
    var employee models.Employee
    if err := c.ShouldBindJSON(&employee); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Create a new employee in the database
    if err := config.DB.Create(&employee).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create employee"})
        return
    }

    // Return the created employee as a response
    c.JSON(http.StatusCreated, employee)
}


func GetEmployees(c *gin.Context) {
    var employees []models.Employee

    if err := config.DB.Find(&employees).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch employees"})
        return
    }

    c.JSON(http.StatusOK, employees)
}


func GetEmployeeDetails(c *gin.Context) {
    userID := c.Param("user_id")
    var employee models.Employee

    if err := config.DB.Where("user_id = ?", userID).First(&employee).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
        return
    }

    c.JSON(http.StatusOK, employee)
}
func DeleteEmployee(c *gin.Context) {
    userID := c.Param("user_id")

    if err := config.DB.Where("user_id = ?", userID).Delete(&models.Employee{}).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete employee"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Employee deleted successfully"})
}
