package routes

import (
	"skillshare-api/controller"
	"skillshare-api/middleware"
	"skillshare-api/repository"
	"skillshare-api/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRoutes(e *echo.Echo, db *gorm.DB) {

	userRepo := repository.NewUserRepository(db)
	classRepo := repository.NewClassRepository(db)
	categoryRepo := repository.NewCategoryRepository(db)
	enrollmentRepo := repository.NewEnrollmentRepository(db)

	
	userService := service.NewUserService(userRepo)
	classService := service.NewClassService(classRepo, userRepo, categoryRepo, enrollmentRepo)

	
	userController := controller.NewUserController(userService)
	classController := controller.NewClassController(classService)
	categoryController := controller.NewCategoryController(db)

	public := e.Group("/api/public")

	public.POST("/register", userController.RegisterUser)
	public.POST("/login", userController.LoginUser)

	public.GET("/classes", classController.GetAllClasses)
	public.GET("/classes/:id", classController.GetClassByID)

	public.GET("/categories", categoryController.GetAllCategories)
	public.GET("/categories/:id", categoryController.GetCategoryByID)

	
	protected := e.Group("/api")
	protected.Use(middleware.JWTMiddleware()) 

	
	protected.GET("/users/:id", userController.GetUserByID)
	protected.PUT("/users/:id", userController.UpdateUser)
	protected.DELETE("/users/:id", userController.DeleteUser)

	
	protected.POST("/classes", classController.CreateClass)
	protected.PUT("/classes/:id", classController.UpdateClass)
	protected.DELETE("/classes/:id", classController.DeleteClass)

	protected.POST("/classes/:class_id/enroll", classController.EnrollInClass)
	protected.GET("/enrollments", classController.GetUserEnrollments)
	protected.DELETE("/classes/:class_id/unenroll", classController.UnenrollFromClass)


	protected.POST("/categories", categoryController.CreateCategory)
	protected.PUT("/categories/:id", categoryController.UpdateCategory)
	protected.DELETE("/categories/:id", categoryController.DeleteCategory)
}
