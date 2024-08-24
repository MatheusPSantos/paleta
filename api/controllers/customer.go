package controllers

import (
	"fmt"
	"net/http"
	"paleta-api/database"
	"paleta-api/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateCustomer(c *gin.Context) {
	var customer models.User
	var auth models.Authentication

	var input struct {
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// extrai a senha do json
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password not provided"})
		return
	}

	if customer.IsSeller {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": "Should be customer"})
		return
	}

	if err := models.ValidateDataUser(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// criptografa senha
	hashedPassword, err := models.HashPassword(input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	if err := database.DB.Create(&customer).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Criar o registro de autenticação associado
	auth.UserID = customer.ID
	auth.Password = hashedPassword

	if err := database.DB.Create(&auth).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create authentication record"})
		return
	}

	c.JSON(http.StatusCreated, customer)
}

func GetCustomers(c *gin.Context) {
	var customers []models.User
	database.DB.Where("is_seller = ?", false).Find(&customers)
	c.JSON(http.StatusOK, customers)
}

func GetCustomerById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var customer models.User

	if err := database.DB.Where("is_seller = ?", false).
		First(&customer, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}

	c.JSON(http.StatusOK, customer)
}

func UpdateCustomer(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var customer models.User
	if err := database.DB.First(&customer, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}

	if customer.IsSeller {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": "This user is an seller"})
		return
	}

	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Save(&customer)
	c.JSON(http.StatusOK, customer)
}

func DeleteCustomer(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var customer models.User

	if err := database.DB.First(&customer, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}

	if customer.IsSeller {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": "This user a seller"})
		return
	}

	if err := database.DB.Delete(&models.User{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Customer deleted"})
}
