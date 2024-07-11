package controllers

import (
	"log"
	"net/http"
	"paleta-api/database"
	"paleta-api/models"

	"github.com/gin-gonic/gin"
)

func CreateLoyaltyCampaing(c *gin.Context) {
	log.Print("Creating loyalty campaing")
	var new_campaing models.LoyaltyCampaing
	var seller models.User

	if err := c.ShouldBindJSON(&new_campaing); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Print("Checking if the owner is and seller.")

	if err := database.DB.Where("is_seller = ?", true).First(&seller, new_campaing.OwnerId).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "The owner of an campaing should be an valid seller user."})
		return
	}

	if err := models.ValidateDataLoyaltyCampaing(&new_campaing); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Create(&new_campaing).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, new_campaing)
}
