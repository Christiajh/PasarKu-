package routes

import (
	"skillshare-api/controller"
	"skillshare-api/middleware"
	"skillshare-api/repository"
	"skillshare-api/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// InitRoutes initializes all API routes
func InitRoutes(e *echo.Echo, db *gorm.DB) {
	// Initialize repositories
	userRepo := repository.NewUserRepository(db)
	classRepo := repository.NewClassRepository(db)
	categoryRepo := repository.NewCategoryRepository(db)
	enrollmentRepo := repository.NewEnrollmentRepository(db)

	// Initialize services
	userService := service.NewUserService(userRepo)
	classService := service.NewClassService(classRepo, userRepo, categoryRepo, enrollmentRepo)

	// Initialize controllers
	userController := controller.NewUserController(userService)
	classController := controller.NewClassController(classService)
	categoryController := controller.NewCategoryController(db) // Category controller directly uses repo

	// Public routes (no authentication required)
	e.POST("/register", userController.RegisterUser)
	e.POST("/login", userController.LoginUser)
	e.GET("/classes", classController.GetAllClasses)
	e.GET("/classes/:id", classController.GetClassByID)
	e.GET("/categories", categoryController.GetAllCategories)
	e.GET("/categories/:id", categoryController.GetCategoryByID)

	// Authenticated routes (JWT required)
	r := e.Group("/api")
	r.Use(middleware.JWTMiddleware())

	// User routes
	r.GET("/users/:id", userController.GetUserByID)
	r.PUT("/users/:id", userController.UpdateUser)
	r.DELETE("/users/:id", userController.DeleteUser)

	// Class routes
	r.POST("/classes", classController.CreateClass)
	r.PUT("/classes/:id", classController.UpdateClass)
	r.DELETE("/classes/:id", classController.DeleteClass)
	r.POST("/classes/:class_id/enroll", classController.EnrollInClass)
	r.GET("/enrollments", classController.GetUserEnrollments)
	r.DELETE("/classes/:class_id/unenroll", classController.UnenrollFromClass)


	// Category routes (consider if these should be admin-only or authenticated)
	// For simplicity, making them authenticated for now, but in a real app,
	// only admins might create/update/delete categories.
	r.POST("/categories", categoryController.CreateCategory)
	r.PUT("/categories/:id", categoryController.UpdateCategory)
	r.DELETE("/categories/:id", categoryController.DeleteCategory)
}