package controller

import (
	"net/http"
	"skillshare-api/helper"
	"skillshare-api/model"
	"skillshare-api/service"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{userService: userService}
}

// RegisterUser handles user registration
func (uc *UserController) RegisterUser(c echo.Context) error {
	var user model.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid request payload"})
	}

	// Password will be stored as plaintext (NOT RECOMMENDED FOR PRODUCTION)
	createdUser, err := uc.userService.RegisterUser(&user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": err.Error()})
	}

	// PENTING: Kosongkan password sebelum mengirimkannya dalam respons JSON
	createdUser.Password = ""

	return c.JSON(http.StatusCreated, echo.Map{
		"message": "User registered successfully",
		"user":    createdUser,
	})
}

// LoginUser handles user login without password hashing verification
// LoginUser handles user login without password hashing verification
func (uc *UserController) LoginUser(c echo.Context) error {
	var loginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.Bind(&loginRequest); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid request payload"})
	}

	user, err := uc.userService.LoginUser(loginRequest.Email, loginRequest.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"message": "Invalid credentials"})
	}

	// ✅ Generate a secure JWT token
	token, err := helper.GenerateJWT(user.ID, user.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Failed to generate token"})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Login successful",
		"token":   token,
	})
}


// GetUserByID retrieves a user by ID
func (uc *UserController) GetUserByID(c echo.Context) error {
	claims, ok := c.Get("user").(*model.JwtCustomClaims)
	if !ok {
		return c.JSON(http.StatusUnauthorized, echo.Map{"message": "Unauthorized: Invalid token claims"})
	}
	loggedInUserID := claims.UserID

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid user ID"})
	}

	if uint(id) != loggedInUserID {
		return c.JSON(http.StatusForbidden, echo.Map{"message": "Forbidden: You can only view your own profile"})
	}

	user, err := uc.userService.GetUserByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"message": err.Error()})
	}

	// PENTING: Kosongkan password sebelum mengirimkannya dalam respons JSON
	user.Password = ""

	return c.JSON(http.StatusOK, echo.Map{
		"message": "User retrieved successfully",
		"user":    user,
	})
}

// UpdateUser updates an existing user
func (uc *UserController) UpdateUser(c echo.Context) error {
	claims, ok := c.Get("user").(*model.JwtCustomClaims)
	if !ok {
		return c.JSON(http.StatusUnauthorized, echo.Map{"message": "Unauthorized: Invalid token claims"})
	}
	loggedInUserID := claims.UserID

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid user ID"})
	}

	if uint(id) != loggedInUserID {
		return c.JSON(http.StatusForbidden, echo.Map{"message": "Forbidden: You can only update your own profile"})
	}

	var updatedUser model.User
	if err := c.Bind(&updatedUser); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid request payload"})
	}
	updatedUser.ID = uint(id)

	user, err := uc.userService.UpdateUser(&updatedUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": err.Error()})
	}

	// PENTING: Kosongkan password sebelum mengirimkannya dalam respons JSON
	user.Password = ""

	return c.JSON(http.StatusOK, echo.Map{
		"message": "User updated successfully",
		"user":    user,
	})
}

// DeleteUser deletes a user
func (uc *UserController) DeleteUser(c echo.Context) error {
	claims, ok := c.Get("user").(*model.JwtCustomClaims)
	if !ok {
		return c.JSON(http.StatusUnauthorized, echo.Map{"message": "Unauthorized: Invalid token claims"})
	}
	loggedInUserID := claims.UserID

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid user ID"})
	}

	if uint(id) != loggedInUserID {
		return c.JSON(http.StatusForbidden, echo.Map{"message": "Forbidden: You can only delete your own profile"})
	}

	err = uc.userService.DeleteUser(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "User deleted successfully"})
}