package main

import (
	"log"
	"net/http" // ADD THIS LINE
	"os"
	"skillshare-api/config"
	"skillshare-api/routes"
	"time"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// 1. Muat environment variables dari file .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// 2. Inisialisasi koneksi database
	db := config.InitDB()
	if db == nil {
		log.Fatal("Failed to connect to database")
	}

	// 3. Inisialisasi Echo framework
	e := echo.New()

	// 4. Tambahkan middleware Echo bawaan
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
        AllowOrigins: []string{"*"}, // Ganti dengan domain frontend Anda di produksi
        AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE, echo.OPTIONS},
        AllowHeaders: []string{"Content-Type", "Authorization"},
    }))

	// 5. Inisialisasi semua rute API
	routes.InitRoutes(e, db)

	// 6. Dapatkan port dari environment variable atau gunakan default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Port default jika PORT tidak diatur di .env
	}

	// 7. Mulai server Echo
	log.Printf("🚀 Server listening on :%s", port)
	s := &http.Server{ // Now 'http' will be defined
		Addr:         ":" + port,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		MaxHeaderBytes: 1 << 20, // 1MB
	}
	e.Logger.Fatal(e.StartServer(s))
}