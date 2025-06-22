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
	// ========== 🧱 REPOSITORIES ==========
	userRepo := repository.NewUserRepository(db)
	classRepo := repository.NewClassRepository(db)
	categoryRepo := repository.NewCategoryRepository(db)
	enrollmentRepo := repository.NewEnrollmentRepository(db)

	// ========== ⚙️ SERVICES ==========
	userService := service.NewUserService(userRepo)
	classService := service.NewClassService(classRepo, userRepo, categoryRepo, enrollmentRepo)

	// ========== 🧠 CONTROLLERS ==========
	userController := controller.NewUserController(userService)
	classController := controller.NewClassController(classService)
	categoryController := controller.NewCategoryController(db)

	// ========== 🔓 PUBLIC ROUTES ==========
	public := e.Group("/api/public")

	public.POST("/register", userController.RegisterUser)
	public.POST("/login", userController.LoginUser)

	public.GET("/classes", classController.GetAllClasses)
	public.GET("/classes/:id", classController.GetClassByID)

	public.GET("/categories", categoryController.GetAllCategories)
	public.GET("/categories/:id", categoryController.GetCategoryByID)

	// ========== 🔐 PROTECTED ROUTES ==========
	protected := e.Group("/api")
	protected.Use(middleware.JWTMiddleware()) // ← middleware custom

	// === 👤 User Routes ===
	protected.GET("/users/:id", userController.GetUserByID)
	protected.PUT("/users/:id", userController.UpdateUser)
	protected.DELETE("/users/:id", userController.DeleteUser)

	// === 📚 Class Routes ===
	protected.POST("/classes", classController.CreateClass)
	protected.PUT("/classes/:id", classController.UpdateClass)
	protected.DELETE("/classes/:id", classController.DeleteClass)

	protected.POST("/classes/:class_id/enroll", classController.EnrollInClass)
	protected.GET("/enrollments", classController.GetUserEnrollments)
	protected.DELETE("/classes/:class_id/unenroll", classController.UnenrollFromClass)

	// === 🏷️ Category Routes ===
	protected.POST("/categories", categoryController.CreateCategory)
	protected.PUT("/categories/:id", categoryController.UpdateCategory)
	protected.DELETE("/categories/:id", categoryController.DeleteCategory)
}
