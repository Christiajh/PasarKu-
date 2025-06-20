package main

import (
	"github.com/gin-gonic/gin"
	"pasarku/database"
	"pasarku/routes"
)

func main() {
	database.Init()
	r := gin.Default()
	routes.SetupRoutes(r)
	r.Run(":8080")
}
