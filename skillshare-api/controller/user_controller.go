package controller

import (
	"net/http"
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

	registeredUser, err := uc.userService.RegisterUser(&user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": err.Error()})
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"message": "User registered successfully",
		"user":    registeredUser,
	})
}

// LoginUser handles user login
func (uc *UserController) LoginUser(c echo.Context) error {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid request payload"})
	}

	token, err := uc.userService.LoginUser(req.Email, req.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Login successful",
		"token":   token,
	})
}

// GetUserByID retrieves a user by ID
func (uc *UserController) GetUserByID(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid user ID"})
	}

	user, err := uc.userService.GetUserByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "User retrieved successfully",
		"user":    user,
	})
}

// UpdateUser updates an existing user
func (uc *UserController) UpdateUser(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid user ID"})
	}

	var updatedUser model.User
	if err := c.Bind(&updatedUser); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid request payload"})
	}

	updatedUser.ID = uint(id) // Ensure the ID from the URL is used

	user, err := uc.userService.UpdateUser(&updatedUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "User updated successfully",
		"user":    user,
	})
}

// DeleteUser deletes a user by ID
func (uc *UserController) DeleteUser(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid user ID"})
	}

	err = uc.userService.DeleteUser(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "User deleted successfully"})
}