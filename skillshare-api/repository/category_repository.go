package repository

import (
	"errors"
	"skillshare-api/model" // Make sure this import is correct

	"gorm.io/gorm"
)

// CategoryRepository defines the methods for interacting with Category data.
type CategoryRepository struct {
	db *gorm.DB
}

// NewCategoryRepository creates a new instance of CategoryRepository.
func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

// Create creates a new category in the database.
func (r *CategoryRepository) Create(category *model.Category) error {
	return r.db.Create(category).Error
}

// FindAll retrieves all categories.
func (r *CategoryRepository) FindAll() ([]model.Category, error) {
	var categories []model.Category
	err := r.db.Find(&categories).Error
	return categories, err
}

// FindByID finds a category by ID.
func (r *CategoryRepository) FindByID(id uint) (*model.Category, error) {
	var category model.Category
	err := r.db.First(&category, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("category not found")
		}
		return nil, err
	}
	return &category, nil
}

// Update updates an existing category in the database.
func (r *CategoryRepository) Update(category *model.Category) error {
	result := r.db.Save(category)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("category not found or no changes made")
	}
	return nil
}

// Delete deletes a category from the database.
func (r *CategoryRepository) Delete(id uint) error {
	result := r.db.Delete(&model.Category{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("category not found")
	}
	return nil
}