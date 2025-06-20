package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pasarku/model"
	"pasarku/database"
)

func CreateProduct(c *gin.Context) {
	var product model.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&product)
	c.JSON(http.StatusCreated, product)
}

func GetProducts(c *gin.Context) {
	var products []model.Product
	database.DB.Find(&products)
	c.JSON(http.StatusOK, products)
}
