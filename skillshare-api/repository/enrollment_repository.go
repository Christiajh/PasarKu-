package repository

import (
	"errors"
	"skillshare-api/model"

	"gorm.io/gorm"
)

type EnrollmentRepository struct {
	db *gorm.DB
}

func NewEnrollmentRepository(db *gorm.DB) *EnrollmentRepository {
	return &EnrollmentRepository{db: db}
}

// Create creates a new enrollment
func (r *EnrollmentRepository) Create(enrollment *model.Enrollment) error {
	return r.db.Create(enrollment).Error
}

// FindByUserIDAndClassID finds an enrollment by user and class ID
func (r *EnrollmentRepository) FindByUserIDAndClassID(userID, classID uint) (*model.Enrollment, error) {
	var enrollment model.Enrollment
	err := r.db.Where("user_id = ? AND class_id = ?", userID, classID).First(&enrollment).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("enrollment not found")
		}
		return nil, err
	}
	return &enrollment, nil
}

// FindAllByUserID retrieves all enrollments for a specific user, with associated class details
func (r *EnrollmentRepository) FindAllByUserID(userID uint) ([]model.Enrollment, error) {
	var enrollments []model.Enrollment
	err := r.db.Preload("Class").Preload("Class.User").Preload("Class.Category").Where("user_id = ?", userID).Find(&enrollments).Error
	return enrollments, err
}

// Delete deletes an enrollment
func (r *EnrollmentRepository) Delete(id uint) error {
	result := r.db.Delete(&model.Enrollment{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("enrollment not found")
	}
	return nil
}

// DeleteByUserIDAndClassID deletes an enrollment by user and class ID
func (r *EnrollmentRepository) DeleteByUserIDAndClassID(userID, classID uint) error {
	result := r.db.Where("user_id = ? AND class_id = ?", userID, classID).Delete(&model.Enrollment{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("enrollment not found")
	}
	return nil
}