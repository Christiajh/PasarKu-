package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pasarku/model"
	"pasarku/database"
)

func CreateOrder(c *gin.Context) {
	var order model.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&order)
	c.JSON(http.StatusCreated, order)
}
