package controllers

import (
	"net/http"
	"paleta-api/database"
	"paleta-api/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateAddress(c *gin.Context) {
	var address models.Address
	if err := c.ShouldBindJSON(&address); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.ValidateDataAddress(&address); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	if err := database.DB.Create(&address).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, address)
}

func GetUserAddress(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var address  []models.Address
	if err := database.DB.Find(&address, "user_id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Address not found"})
		println(err.Error())
		return
	}
	c.JSON(http.StatusOK, address)
}

func UpdateUserAddress(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("id"))
	addressId, _ := strconv.Atoi(c.Param("addressId"))

	var address models.Address

	if err := database.DB.Where("user_id = ?", userId).First(&address, addressId).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Address not found"})
		return
	}

	if err := c.ShouldBindJSON(&address); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Save(&address)
	c.JSON(http.StatusOK, address)
}

func DeleteUserAddress(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("id"))
	addressId, _ := strconv.Atoi(c.Param("addressId"))
	var address models.Address
	if err := database.DB.Where("user_id = ?", userId).First(&address, addressId).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Addres not found."})
		return
	}
	if err := database.DB.Where("user_id = ?", userId).Delete(&models.Address{}, addressId).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Addreess not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Address deleted."})
}
