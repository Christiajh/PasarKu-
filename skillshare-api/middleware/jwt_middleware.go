package middleware

import (
	"os"
	"skillshare-api/model"

	"github.com/golang-jwt/jwt/v5"       // Still need this for jwt.Claims and RegisteredClaims
	"github.com/labstack/echo/v4"
	echojwt "github.com/labstack/echo-jwt/v4" // ALIAS THIS PACKAGE for clarity and correct usage
)

// JWTSecret returns the JWT secret key from environment variables
func JWTSecret() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "your_jwt_secret_key" // Fallback for local development if .env is missing
	}
	return secret
}

// JWTMiddleware creates an Echo JWT middleware instance
func JWTMiddleware() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(model.JwtCustomClaims)
		},
		SigningKey: []byte(JWTSecret()),
		ContextKey: "user", // The key to store the JWT token in the context
		TokenLookup: "header:Authorization:Bearer",
		// REMOVED: AuthScheme:  "Bearer", // This field is likely no longer supported in the latest echo-jwt/v4
		KeyFunc: func(token *jwt.Token) (interface{}, error) {
			// Check the signing method
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(JWTSecret()), nil
		},
	})
}