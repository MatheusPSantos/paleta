package controllers

import (
	"log"
	"net/http"
	"paleta-api/database"
	"paleta-api/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetSellers retrieves all sellers
func GetSellers(c *gin.Context) {
	log.Print("Getting sellers")
	var sellers []models.Seller
	database.DB.Find(&sellers)
	c.JSON(http.StatusOK, sellers)
}

// GetSeller retrieves a seller by ID
func GetSeller(c *gin.Context) {
	log.Print(">>>>>>>>>>>")
	id, _ := strconv.Atoi(c.Param("id"))
	var seller models.Seller
	if err := database.DB.First(&seller, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Seller not found"})
		return
	}
	c.JSON(http.StatusOK, seller)
}

// CreateSeller creates a new seller
func CreateSeller(c *gin.Context) {
	var seller models.Seller
	if err := c.ShouldBindJSON(&seller); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&seller)
	c.JSON(http.StatusOK, seller)
}

// UpdateSeller updates an existing seller
func UpdateSeller(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var seller models.Seller
	if err := database.DB.First(&seller, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Seller not found"})
		return
	}

	if err := c.ShouldBindJSON(&seller); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Save(&seller)
	c.JSON(http.StatusOK, seller)
}

// DeleteSeller deletes a seller
func DeleteSeller(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := database.DB.Delete(&models.Seller{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Seller not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Seller deleted"})
}
