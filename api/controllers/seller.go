package controllers

import (
	"log"
	"net/http"
	"paleta-api/database"
	"paleta-api/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetSellers(c *gin.Context) {
	log.Print("Getting sellers")
	var sellers []models.Seller
	database.DB.Find(&sellers)
	c.JSON(http.StatusOK, sellers)
}

func GetSeller(c *gin.Context) {
	log.Print("[Seller] - getting seller by id")
	id, _ := strconv.Atoi(c.Param("id"))
	var seller models.Seller
	if err := database.DB.Select("id", "created_at", "updated_at", "deleted_at", "name", "email", "phone", "cpf_cnpj").First(&seller, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Seller not found"})
		return
	}
	c.JSON(http.StatusOK, seller)
}

func CreateSeller(c *gin.Context) {
	var seller models.Seller
	if err := c.ShouldBindJSON(&seller); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.ValidateDataSeller(&seller); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	if err := database.DB.Create(&seller).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, seller)
}

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

func DeleteSeller(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := database.DB.Delete(&models.Seller{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Seller not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Seller deleted"})
}

func GetSellerAddress(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var address models.Address
	if err := database.DB.First(&address, "seller_id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Address not found"})
		println(err.Error())
		return
	}
	c.JSON(http.StatusOK, address)
}
