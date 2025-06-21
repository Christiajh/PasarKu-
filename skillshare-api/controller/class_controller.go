package controller

import (
	"net/http"
	"skillshare-api/model"
	"skillshare-api/service"
	"strconv"

	"github.com/labstack/echo/v4"
	// No longer need "github.com/golang-jwt/jwt" here directly for the assertion,
	// as you're asserting to your model's claims, which uses jwt.RegisteredClaims internally.
)

type ClassController struct {
	classService *service.ClassService
}

func NewClassController(classService *service.ClassService) *ClassController {
	return &ClassController{classService: classService}
}

// CreateClass handles creating a new class
func (cc *ClassController) CreateClass(c echo.Context) error {
	// CORRECTED: Assert directly to *model.JwtCustomClaims
	claims, ok := c.Get("user").(*model.JwtCustomClaims)
	if !ok {
		return c.JSON(http.StatusUnauthorized, echo.Map{"message": "Unauthorized: Invalid token claims"})
	}
	userID := claims.UserID

	var class model.Class
	if err := c.Bind(&class); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid request payload"})
	}
	class.UserID = userID // Set the user ID from the authenticated token

	createdClass, err := cc.classService.CreateClass(&class)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": err.Error()})
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"message": "Class created successfully",
		"class":   createdClass,
	})
}

// GetAllClasses retrieves all classes (no authentication needed here)
func (cc *ClassController) GetAllClasses(c echo.Context) error {
	classes, err := cc.classService.GetAllClasses()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Classes retrieved successfully",
		"classes": classes,
	})
}

// GetClassByID retrieves a class by ID (no authentication needed here)
func (cc *ClassController) GetClassByID(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid class ID"})
	}

	class, err := cc.classService.GetClassByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Class retrieved successfully",
		"class":   class,
	})
}

// UpdateClass updates an existing class
func (cc *ClassController) UpdateClass(c echo.Context) error {
	// CORRECTED: Assert directly to *model.JwtCustomClaims
	claims, ok := c.Get("user").(*model.JwtCustomClaims)
	if !ok {
		return c.JSON(http.StatusUnauthorized, echo.Map{"message": "Unauthorized: Invalid token claims"})
	}
	userID := claims.UserID

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid class ID"})
	}

	var updatedClass model.Class
	if err := c.Bind(&updatedClass); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid request payload"})
	}
	updatedClass.ID = uint(id)

	class, err := cc.classService.UpdateClass(&updatedClass, userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Class updated successfully",
		"class":   class,
	})
}

// DeleteClass deletes a class
func (cc *ClassController) DeleteClass(c echo.Context) error {
	// CORRECTED: Assert directly to *model.JwtCustomClaims
	claims, ok := c.Get("user").(*model.JwtCustomClaims)
	if !ok {
		return c.JSON(http.StatusUnauthorized, echo.Map{"message": "Unauthorized: Invalid token claims"})
	}
	userID := claims.UserID

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid class ID"})
	}

	err = cc.classService.DeleteClass(uint(id), userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "Class deleted successfully"})
}

// EnrollInClass handles user enrollment in a class
func (cc *ClassController) EnrollInClass(c echo.Context) error {
	// CORRECTED: Assert directly to *model.JwtCustomClaims
	claims, ok := c.Get("user").(*model.JwtCustomClaims)
	if !ok {
		return c.JSON(http.StatusUnauthorized, echo.Map{"message": "Unauthorized: Invalid token claims"})
	}
	userID := claims.UserID

	classID, err := strconv.ParseUint(c.Param("class_id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid class ID"})
	}

	enrollment, err := cc.classService.EnrollInClass(userID, uint(classID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": err.Error()})
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"message":    "Successfully enrolled in class",
		"enrollment": enrollment,
	})
}

// GetUserEnrollments retrieves all classes a user is enrolled in
func (cc *ClassController) GetUserEnrollments(c echo.Context) error {
	// CORRECTED: Assert directly to *model.JwtCustomClaims
	claims, ok := c.Get("user").(*model.JwtCustomClaims)
	if !ok {
		return c.JSON(http.StatusUnauthorized, echo.Map{"message": "Unauthorized: Invalid token claims"})
	}
	userID := claims.UserID

	enrollments, err := cc.classService.GetUserEnrollments(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message":     "User enrollments retrieved successfully",
		"enrollments": enrollments,
	})
}

// UnenrollFromClass handles unenrollment from a class
func (cc *ClassController) UnenrollFromClass(c echo.Context) error {
	// CORRECTED: Assert directly to *model.JwtCustomClaims
	claims, ok := c.Get("user").(*model.JwtCustomClaims)
	if !ok {
		return c.JSON(http.StatusUnauthorized, echo.Map{"message": "Unauthorized: Invalid token claims"})
	}
	userID := claims.UserID

	classID, err := strconv.ParseUint(c.Param("class_id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid class ID"})
	}

	err = cc.classService.UnenrollFromClass(userID, uint(classID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "Successfully unenrolled from class"})
}