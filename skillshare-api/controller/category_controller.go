package controller

import (
	"net/http"
	"skillshare-api/model"
	"skillshare-api/repository"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type CategoryController struct {
	categoryRepo *repository.CategoryRepository
}

func NewCategoryController(db *gorm.DB) *CategoryController {
	return &CategoryController{categoryRepo: repository.NewCategoryRepository(db)}
}

// CreateCategory handles creating a new category
func (cc *CategoryController) CreateCategory(c echo.Context) error {
	var category model.Category
	if err := c.Bind(&category); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid request payload"})
	}

	if err := cc.categoryRepo.Create(&category); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Failed to create category"})
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"message":  "Category created successfully",
		"category": category,
	})
}

// GetAllCategories retrieves all categories
func (cc *CategoryController) GetAllCategories(c echo.Context) error {
	categories, err := cc.categoryRepo.FindAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Failed to retrieve categories"})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message":    "Categories retrieved successfully",
		"categories": categories,
	})
}

// GetCategoryByID retrieves a category by ID
func (cc *CategoryController) GetCategoryByID(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid category ID"})
	}

	category, err := cc.categoryRepo.FindByID(uint(id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, echo.Map{"message": "Category not found"})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Failed to retrieve category"})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message":  "Category retrieved successfully",
		"category": category,
	})
}

// UpdateCategory updates an existing category
func (cc *CategoryController) UpdateCategory(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid category ID"})
	}

	var updatedCategory model.Category
	if err := c.Bind(&updatedCategory); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid request payload"})
	}
	updatedCategory.ID = uint(id)

	if err := cc.categoryRepo.Update(&updatedCategory); err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, echo.Map{"message": "Category not found"})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Failed to update category"})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message":  "Category updated successfully",
		"category": updatedCategory,
	})
}

// DeleteCategory deletes a category
func (cc *CategoryController) DeleteCategory(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid category ID"})
	}

	if err := cc.categoryRepo.Delete(uint(id)); err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, echo.Map{"message": "Category not found"})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Failed to delete category"})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "Category deleted successfully"})
}