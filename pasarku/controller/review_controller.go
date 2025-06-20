package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pasarku/model"
	"pasarku/database"
)

func CreateReview(c *gin.Context) {
	var review model.Review
	if err := c.ShouldBindJSON(&review); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&review)
	c.JSON(http.StatusCreated, review)
}
