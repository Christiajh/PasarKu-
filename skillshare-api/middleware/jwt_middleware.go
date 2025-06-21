package middleware

import (
	"fmt"
	"net/http"
	"time"

	"skillshare-api/helper"
	"skillshare-api/model"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	echojwt "github.com/labstack/echo-jwt/v4"
)

// JWTMiddleware returns Echo middleware that verifies JWT tokens
func JWTMiddleware() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		// Gunakan custom claims
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(model.JwtCustomClaims)
		},

		// Konfigurasi algoritma dan secret
		SigningKey:    []byte(helper.JWTSecret()),
		SigningMethod: "HS256",

		// Token diambil dari header Authorization: Bearer <token>
		TokenLookup: "header:Authorization:Bearer",

		// Token disimpan dalam context dengan key "user"
		ContextKey: "user",

		// Handler jika token invalid/expired
		ErrorHandler: func(c echo.Context, err error) error {
			authHeader := c.Request().Header.Get("Authorization")

			// Debug Logging
			fmt.Println("🔒 JWT Error Handler Triggered")
			fmt.Printf("📥 JWT Auth Header: %q\n", authHeader)
			fmt.Printf("📥 Header Length: %d\n", len(authHeader))
			fmt.Println("🔐 JWT Secret Used:", helper.JWTSecret())
			fmt.Println("🕒 Server Time:", time.Now().Format(time.RFC3339))
			fmt.Printf("❌ JWT Middleware Error: %v\n", err)

			return c.JSON(http.StatusUnauthorized, echo.Map{
				"message": "invalid or expired jwt",
				"detail":  err.Error(),
			})
		},
	})
}
