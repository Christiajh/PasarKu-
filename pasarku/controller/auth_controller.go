package controller

import (
	"net/http"
	"pasarku/helper"
	"pasarku/model"
	"pasarku/structs"

	"github.com/gin-gonic/gin"
)

// Register godoc - membuat user baru
func Register(c *gin.Context) {
	var input structs.RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		helper.ErrorResponse(c, http.StatusBadRequest, "Invalid input", err.Error())
		return
	}

	// Cek apakah email sudah digunakan
	if _, err := model.GetUserByEmail(input.Email); err == nil {
		helper.ErrorResponse(c, http.StatusConflict, "Email already registered", nil)
		return
	}

	hashedPassword, err := helper.HashPassword(input.Password)
	if err != nil {
		helper.ErrorResponse(c, http.StatusInternalServerError, "Failed to hash password", err.Error())
		return
	}

	user := model.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: hashedPassword,
		Role:     "user", // default role
	}

	if err := model.CreateUser(&user); err != nil {
		helper.ErrorResponse(c, http.StatusInternalServerError, "Failed to register user", err.Error())
		return
	}

	helper.SuccessResponse(c, http.StatusCreated, "User registered successfully", gin.H{
		"id":    user.ID,
		"name":  user.Name,
		"email": user.Email,
		"role":  user.Role,
	})
}

// Login godoc - autentikasi user & kembalikan token
func Login(c *gin.Context) {
	var input structs.LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		helper.ErrorResponse(c, http.StatusBadRequest, "Invalid input", err.Error())
		return
	}

	user, err := model.GetUserByEmail(input.Email)
	if err != nil {
		helper.ErrorResponse(c, http.StatusUnauthorized, "Email not found", nil)
		return
	}

	if err := helper.CheckPassword(user.Password, input.Password); err != nil {
		helper.ErrorResponse(c, http.StatusUnauthorized, "Wrong password", nil)
		return
	}

	// Buat token JWT
	token, err := helper.GenerateToken(user.ID, user.Role)
	if err != nil {
		helper.ErrorResponse(c, http.StatusInternalServerError, "Failed to generate token", err.Error())
		return
	}

	helper.SuccessResponse(c, http.StatusOK, "Login successful", gin.H{
		"token": token,
	})
}

// GetProfile godoc - ambil info user dari token
func GetProfile(c *gin.Context) {
	userID := c.MustGet("userID").(uint)

	user, err := model.GetUserByID(userID)
	if err != nil {
		helper.ErrorResponse(c, http.StatusNotFound, "User not found", nil)
		return
	}

	helper.SuccessResponse(c, http.StatusOK, "User profile retrieved", gin.H{
		"id":    user.ID,
		"name":  user.Name,
		"email": user.Email,
		"role":  user.Role,
	})
}
