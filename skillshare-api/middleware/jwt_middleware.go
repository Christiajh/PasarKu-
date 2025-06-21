package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"
	// "strings" // <-- HAPUS INI, karena kita tidak akan menggunakan strings.TrimSpace di ErrorHandler

	"github.com/labstack/echo/v4"
	echojwt "github.com/labstack/echo-jwt/v4" // Pastikan versi v4
)

// JWTSecret retrieves the JWT secret key from environment variables.
// Ini adalah sumber kebenaran untuk secret key saat token DIVERIFIKASI.
// Pastikan nilainya SAMA PERSIS dengan yang di helper/helper.go
func JWTSecret() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		fmt.Println("⚠️ WARNING: JWT_SECRET not set. Using fallback. DO NOT use in production.")
		secret = "your_super_secure_and_long_consistent_key"
	}
	return secret
}

// JWTMiddleware returns a middleware that validates JWT tokens.
func JWTMiddleware() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		SigningKey:    []byte(JWTSecret()),
		SigningMethod: "HS256",
		ContextKey:    "user",
		TokenLookup:   "header:Authorization:Bearer",
		ErrorHandler: func(c echo.Context, err error) error {
			authHeader := c.Request().Header.Get("Authorization")

			// --- LOGGING YANG SANGAT KRUSIAL UNTUK DEBUGGING ---
			// `%q` akan menampilkan karakter non-printable (seperti spasi ekstra, null byte, dll.)
			// Ini akan menunjukkan persis apa yang diterima server dari Postman.
			fmt.Printf("📥 JWT Middleware menerima Authorization (raw): %q\n", authHeader)
			fmt.Printf("📥 JWT Middleware menerima Authorization (length): %d\n", len(authHeader))
			// --- AKHIR LOGGING KRUSIAL ---

			fmt.Println("🔐 JWT_SECRET yang digunakan:", JWTSecret())
			fmt.Println("🕒 Waktu server:", time.Now().Format(time.RFC3339))
			fmt.Printf("❌ JWT Middleware Error: %v\n", err) // Log error yang sebenarnya

			return echo.NewHTTPError(http.StatusUnauthorized, map[string]interface{}{
				"message": "invalid or expired jwt",
				"detail":  err.Error(),
			})
		},
	})
}