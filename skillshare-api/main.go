package main

import (
	"log"
	"net/http"
	"os"
	"skillshare-api/config"
	"skillshare-api/routes"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	db := config.ConnectDatabase()

	
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE, echo.OPTIONS},
		AllowHeaders: []string{"Content-Type", "Authorization"},
	}))

	
	routes.InitRoutes(e, db)

	
	e.Static("/docs", "docs") 
	e.File("/swagger", "docs/index.html") 

	
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("🚀 Server listening on :%s", port)
	s := &http.Server{
		Addr:           ":" + port,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	e.Logger.Fatal(e.StartServer(s))
}
