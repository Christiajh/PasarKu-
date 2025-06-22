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

// JWTMiddleware provides Bearer-authenticated JWT middleware
func JWTMiddleware() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(model.JwtCustomClaims)
		},
		SigningKey:    []byte(helper.JWTSecret()),
		SigningMethod: "HS256",

		// ✅ Token dengan format: "Authorization: Bearer <token>"
		TokenLookup: "header:Authorization",


		ContextKey: "user",

		ErrorHandler: func(c echo.Context, err error) error {
			authHeader := c.Request().Header.Get("Authorization")

			// Debug log
			fmt.Println("🔒 JWT Middleware Error Triggered")
			fmt.Printf("📥 Authorization Header: %q\n", authHeader)
			fmt.Println("🕒 Server Time:", time.Now().Format(time.RFC3339))
			fmt.Printf("❌ JWT Error: %v\n", err)

			return c.JSON(http.StatusUnauthorized, echo.Map{
				"message": "invalid or expired jwt",
				"detail":  err.Error(),
			})
		},
	})
}
