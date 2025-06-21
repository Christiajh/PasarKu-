package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/golang-jwt/jwt/v5" // Make sure you import jwt/v5 if you're using it for claims
)

// JWTSecret retrieves the JWT secret key from environment variables.
// This function ensures the secret used for signing and verification is consistent.
func JWTSecret() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		fmt.Println("⚠️ WARNING: JWT_SECRET not set. Using fallback. DO NOT use in production.")
		// Using "TEST_SECRET_123" is fine for temporary debugging,
		// but remember to change it to a strong, consistent key for production.
		secret = "TEST_SECRET_123"
	}
	return secret
}

// JWTMiddleware returns a middleware that validates JWT tokens.
func JWTMiddleware() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		// NewClaimsFunc is required when using jwt/v5.
		// You need to replace 'jwt.RegisteredClaims' with your actual claims struct
		// (e.g., 'new(model.JwtCustomClaims)' if you have one defined).
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			// Replace with your actual custom claims struct, e.g.,
			// return new(model.JwtCustomClaims)
			return new(jwt.RegisteredClaims) // Using basic RegisteredClaims as a fallback example
		},
		SigningKey:    []byte(JWTSecret()),
		SigningMethod: "HS256", // Ensure this matches how your tokens are signed
		ContextKey:    "user",

		// --- THE CRITICAL FIX IS HERE ---
		// This tells echo-jwt to look in the Authorization header,
		// expect a "Bearer " prefix, and also look in the "jwt" cookie.
		TokenLookup: "header:Authorization:Bearer ,cookie:jwt",
		// --- END OF FIX ---

		ErrorHandler: func(c echo.Context, err error) error {
			authHeader := c.Request().Header.Get("Authorization")

			// These logs are excellent for debugging and should now show the correctly extracted token
			fmt.Printf("📥 JWT Middleware menerima Authorization (raw): %q\n", authHeader)
			fmt.Printf("📥 JWT Middleware menerima Authorization (length): %d\n", len(authHeader))

			fmt.Println("🔐 JWT_SECRET yang digunakan:", JWTSecret())
			fmt.Println("🕒 Waktu server:", time.Now().Format(time.RFC3339))
			fmt.Printf("❌ JWT Middleware Error: %v\n", err) // Log the actual error

			return echo.NewHTTPError(http.StatusUnauthorized, map[string]interface{}{
				"message": "invalid or expired jwt",
				"detail":  err.Error(),
			})
		},
	})
}