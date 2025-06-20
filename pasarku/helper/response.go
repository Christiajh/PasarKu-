package helper

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// SuccessResponse mengirimkan response sukses standar
func SuccessResponse(c *gin.Context, statusCode int, message string, data interface{}) {
	c.JSON(statusCode, gin.H{
		"success": true,
		"message": message,
		"data":    data,
	})
}

// ErrorResponse mengirimkan response error dengan status dan pesan
func ErrorResponse(c *gin.Context, statusCode int, message string, errors interface{}) {
	c.AbortWithStatusJSON(statusCode, gin.H{
		"success": false,
		"message": message,
		"error":   errors,
	})
}

// BadRequest helper untuk 400
func BadRequest(c *gin.Context, message string) {
	ErrorResponse(c, http.StatusBadRequest, message, nil)
}

// NotFound helper untuk 404
func NotFound(c *gin.Context, message string) {
	ErrorResponse(c, http.StatusNotFound, message, nil)
}

// Unauthorized helper untuk 401
func Unauthorized(c *gin.Context, message string) {
	ErrorResponse(c, http.StatusUnauthorized, message, nil)
}
