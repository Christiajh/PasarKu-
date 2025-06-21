package repository

import (
	"errors"
	"skillshare-api/model"

	"gorm.io/gorm"
)

type ClassRepository struct {
	db *gorm.DB
}

func NewClassRepository(db *gorm.DB) *ClassRepository {
	return &ClassRepository{db: db}
}

// Create creates a new class in the database
func (r *ClassRepository) Create(class *model.Class) error {
	return r.db.Create(class).Error
}

// FindAll retrieves all classes, eagerly loading User and Category
func (r *ClassRepository) FindAll() ([]model.Class, error) {
	var classes []model.Class
	err := r.db.Preload("User").Preload("Category").Find(&classes).Error
	return classes, err
}

// FindByID finds a class by ID, eagerly loading User and Category
func (r *ClassRepository) FindByID(id uint) (*model.Class, error) {
	var class model.Class
	err := r.db.Preload("User").Preload("Category").First(&class, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("class not found")
		}
		return nil, err
	}
	return &class, nil
}

// Update updates an existing class in the database
func (r *ClassRepository) Update(class *model.Class) error {
	result := r.db.Save(class)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("class not found or no changes made")
	}
	return nil
}

// Delete deletes a class from the database
func (r *ClassRepository) Delete(id uint) error {
	result := r.db.Delete(&model.Class{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("class not found")
	}
	return nil
}