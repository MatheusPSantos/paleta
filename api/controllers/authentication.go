package controllers

import (
	"fmt"
	"net/http"
	"paleta-api/database"
	"paleta-api/models"
	"paleta-api/services"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var loginCustomer models.Login
	var user models.User
	var auth models.Authentication

	if err := c.ShouldBindJSON(&loginCustomer); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	if err := models.ValidateDataLogin(&loginCustomer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userIdent := loginCustomer.User

	if err := database.DB.Where("username = ?", userIdent).
		Or("cpf_cnpj = ?", userIdent).
		Or("email = ?", userIdent).
		First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if err := database.DB.Where("user_id = ?", user.ID).First(&auth).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Authentication record not found"})
		return
	}
	fmt.Println(user.ID)
	fmt.Println(loginCustomer.Password)
	fmt.Println(auth)
	fmt.Println(user)
	var pass string = loginCustomer.Password
	
	if !auth.CheckPassword(pass) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}
	// gerar token JWT
	token, err := services.GenerateJWT(user.Email + user.Username + user.Name + user.Cpf_Cnpj)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}
