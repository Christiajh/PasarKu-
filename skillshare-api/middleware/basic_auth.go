package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

// BasicAuth is an example of a basic authentication middleware.
// You might use this for admin routes or specific internal API calls,
// but for most user-facing APIs, JWT is preferred.
func BasicAuth(username, password string) echo.MiddlewareFunc {
	return echoMiddleware.BasicAuth(func(u, p string, c echo.Context) (bool, error) {
		if u == username && p == password {
			return true, nil
		}
		return false, echo.NewHTTPError(http.StatusUnauthorized, "Invalid credentials")
	})
}