package routes

import (
	"github.com/gin-gonic/gin"
	"pasarku/controller"
	"pasarku/middleware"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/api/register", controller.Register)
	r.POST("/api/login", controller.Login)

	api := r.Group("/api")
	api.Use(middleware.JWTAuthMiddleware())
	{
		api.GET("/products", controller.GetProducts)
		api.POST("/products", controller.CreateProduct)
		api.POST("/orders", controller.CreateOrder)
		api.POST("/products/:id/review", controller.CreateReview)
	}
}
