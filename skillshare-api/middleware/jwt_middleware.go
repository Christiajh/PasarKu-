package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"skillshare-api/helper"
	"skillshare-api/model"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

// JWTMiddleware is a custom JWT middleware that parses the token manually
func JWTMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")

			if authHeader == "" {
				return c.JSON(http.StatusUnauthorized, echo.Map{
					"message": "Missing Authorization header",
				})
			}

			// 🛡️ Handle "Bearer <token>" or raw token
			var tokenString string
			if strings.HasPrefix(strings.ToLower(authHeader), "bearer ") {
				tokenString = authHeader[7:] // Strip "Bearer "
			} else {
				tokenString = authHeader
			}

			// 🧾 Parse JWT token
			token, err := jwt.ParseWithClaims(tokenString, &model.JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
				if token.Method.Alg() != jwt.SigningMethodHS256.Alg() {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				return []byte(helper.JWTSecret()), nil
			})

			if err != nil {
				// Log detail untuk debugging
				fmt.Println("🔒 JWT Middleware Error Triggered")
				fmt.Printf("📥 Authorization Header: %q\n", authHeader)
				fmt.Println("🕒 Server Time:", time.Now().Format(time.RFC3339))
				fmt.Printf("❌ JWT Error: %v\n", err)

				return c.JSON(http.StatusUnauthorized, echo.Map{
					"message": "invalid or expired jwt",
					"detail":  err.Error(),
				})
			}

			// ✅ Validasi claims dan simpan ke context
			if claims, ok := token.Claims.(*model.JwtCustomClaims); ok && token.Valid {
				c.Set("user", claims)
				return next(c)
			}

			return c.JSON(http.StatusUnauthorized, echo.Map{
				"message": "invalid token claims",
			})
		}
	}
}
