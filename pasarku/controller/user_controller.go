package controller

import (
	"net/http"
	"pasarku/helper"
	"pasarku/model"
	"pasarku/structs"
	"pasarku/helper/validation"
	"github.com/gin-gonic/gin"
)

// Register handles user registration
func Register(c *gin.Context) {
	var input structs.RegisterInput

	// Bind JSON input
	if err := c.ShouldBindJSON(&input); err != nil {
		helper.ErrorResponse(c, http.StatusBadRequest, "Invalid input", validation.FormatValidationError(err))
		return
	}

	// Hash password
	hashedPassword, err := helper.HashPassword(input.Password)
	if err != nil {
		helper.ErrorResponse(c, http.StatusInternalServerError, "Failed to hash password", nil)
		return
	}

	user := model.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: hashedPassword,
		Role:     "user", // default role
	}

	// Simpan ke DB
	if err := model.CreateUser(&user); err != nil {
		helper.ErrorResponse(c, http.StatusInternalServerError, "Failed to register user", nil)
		return
	}

	helper.SuccessResponse(c, http.StatusCreated, "User registered successfully", gin.H{"user": user})
}

// Login handles user login
func Login(c *gin.Context) {
	var input structs.LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		helper.ErrorResponse(c, http.StatusBadRequest, "Invalid input", validation.FormatValidationError(err))
		return
	}

	user, err := model.GetUserByEmail(input.Email)
	if err != nil {
		helper.ErrorResponse(c, http.StatusUnauthorized, "Invalid credentials", nil)
		return
	}

	// Bandingkan password
	if err := helper.CheckPassword(user.Password, input.Password); err != nil {
		helper.ErrorResponse(c, http.StatusUnauthorized, "Invalid credentials", nil)
		return
	}

	// Generate JWT
	token, err := helper.GenerateToken(user.ID, user.Role)
	if err != nil {
		helper.ErrorResponse(c, http.StatusInternalServerError, "Failed to generate token", nil)
		return
	}

	helper.SuccessResponse(c, http.StatusOK, "Login successful", gin.H{"token": token})
}

// GetProfile returns the profile of the logged-in user
func GetProfile(c *gin.Context) {
	userID := c.MustGet("userID").(uint)

	user, err := model.GetUserByID(userID)
	if err != nil {
		helper.ErrorResponse(c, http.StatusNotFound, "User not found", nil)
		return
	}

	helper.SuccessResponse(c, http.StatusOK, "User profile retrieved", gin.H{"user": user})
}
