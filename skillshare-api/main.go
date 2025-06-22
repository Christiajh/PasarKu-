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
	// 1. Inisialisasi koneksi database (gunakan DATABASE_URL dari environment)
	db := config.ConnectDatabase()

	// 2. Inisialisasi Echo framework
	e := echo.New()

	// 3. Tambahkan middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE, echo.OPTIONS},
		AllowHeaders: []string{"Content-Type", "Authorization"},
	}))

	// 4. Inisialisasi semua rute API
	routes.InitRoutes(e, db)

	// 5. Tambahkan static file Swagger UI (dari folder docs)
	e.Static("/docs", "docs") // Pastikan folder 'docs' berisi index.html, swagger.yaml, dll
	e.File("/swagger", "docs/index.html") // Optional: akses langsung ke /swagger

	// 6. Gunakan PORT dari environment (Railway akan mengisi ini otomatis)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // fallback untuk lokal
	}

	// 7. Jalankan server
	log.Printf("🚀 Server listening on :%s", port)
	s := &http.Server{
		Addr:           ":" + port,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	e.Logger.Fatal(e.StartServer(s))
}
