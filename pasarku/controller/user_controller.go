package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pasarku/model"
	"pasarku/database"
	"pasarku/helper"
)

func Register(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&user)
	c.JSON(http.StatusCreated, user)
}

func Login(c *gin.Context) {
	var input model.LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var user model.User
	database.DB.Where("email = ? AND password = ?", input.Email, input.Password).First(&user)
	if user.ID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}
	token, err := helper.GenerateJWT(user.ID, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Token error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}
