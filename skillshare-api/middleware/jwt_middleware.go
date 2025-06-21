package middleware

import (
	"fmt" // Import fmt for logging warnings/errors
	"net/http" // Import http for HTTP status codes
	"os"
	"skillshare-api/model" // Make sure this path is correct

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	echojwt "github.com/labstack/echo-jwt/v4" // Aliased for clarity
)

// JWTSecret retrieves the JWT secret key from environment variables.
//
// !!! CRITICAL: Ensure the JWT_SECRET environment variable is set
// on Railway to the EXACT SAME value used when signing the tokens.
// If it's not set, this function will use a fallback, which can
// lead to "invalid or expired jwt" errors if the signing key differs.
func JWTSecret() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		// Log a warning if the secret is not set via environment variable.
		// In a production environment, you might want to panic here
		// or return an error to prevent the application from starting
		// with an insecure or mismatched secret.
		fmt.Println("WARNING: JWT_SECRET environment variable is not set. Using a default fallback secret. ENSURE THIS IS INTENTIONAL FOR DEVELOPMENT. FOR PRODUCTION, ALWAYS SET VIA ENV VAR.")
		// This fallback MUST match the secret used in helper/jwt.go for signing
		// if you're relying on it locally or if Railway somehow misses the env var.
		secret = "a_very_secret_and_long_key_for_jwt_validation_12345" // <-- !!! GANTI DENGAN KUNCI RAHASIA ASLI ANDA !!!
	}
	return secret
}

// JWTMiddleware creates an Echo JWT middleware instance.
// This middleware is responsible for validating incoming JWT tokens
// on protected API routes.
func JWTMiddleware() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		// NewClaimsFunc tells the middleware which claims struct to use for parsing.
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(model.JwtCustomClaims) // Return a pointer to your custom claims struct
		},
		// SigningKey provides the key used to verify the JWT's signature.
		// This must match the key used to sign the token by helper.GenerateJWT.
		SigningKey: []byte(JWTSecret()),
		// ContextKey is the key under which the parsed JWT claims will be stored
		// in the Echo context for later retrieval by handlers.
		ContextKey: "user",
		// TokenLookup specifies where the middleware should look for the JWT.
		// "header:Authorization:Bearer" means it looks in the "Authorization"
		// header for a value prefixed with "Bearer ".
		TokenLookup: "header:Authorization:Bearer",
		// KeyFunc is generally not needed if SigningKey is used for HMAC.
		// It's useful for more complex scenarios (e.g., multiple keys, different algorithms).
		// We're omitting it for simplicity and correctness with echo-jwt/v4 unless specifically required.
		//
		// KeyFunc: func(token *jwt.Token) (interface{}, error) {
		// 	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		// 		return nil, jwt.ErrSignatureInvalid // Or a custom error
		// 	}
		// 	return []byte(JWTSecret()), nil
		// },

		// ErrorHandler is an optional function to customize how JWT errors are handled.
		// It's highly recommended for debugging in development.
		ErrorHandler: func(c echo.Context, err error) error {
			// Log the actual underlying error for server-side debugging.
			c.Logger().Errorf("JWT Middleware Error: %v", err)

			// Return a generic unauthorized error to the client for security.
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid or expired jwt")
		},

		// Leeway is an optional configuration to allow for clock skew (time differences).
		// If your server and client/token issuer clocks are slightly out of sync,
		// a small leeway (e.g., 5 seconds or 1 minute) can prevent premature expiration errors.
		// Example: Leeway: time.Duration(1 * time.Minute),
	})
}