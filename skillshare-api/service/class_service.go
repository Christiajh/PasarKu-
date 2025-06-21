package service

import (
	"errors"
	"skillshare-api/model"
	"skillshare-api/repository"
)

type ClassService struct {
	classRepo      *repository.ClassRepository
	userRepo       *repository.UserRepository
	categoryRepo   *repository.CategoryRepository
	enrollmentRepo *repository.EnrollmentRepository
}

func NewClassService(classRepo *repository.ClassRepository, userRepo *repository.UserRepository, categoryRepo *repository.CategoryRepository, enrollmentRepo *repository.EnrollmentRepository) *ClassService {
	return &ClassService{
		classRepo:      classRepo,
		userRepo:       userRepo,
		categoryRepo:   categoryRepo,
		enrollmentRepo: enrollmentRepo,
	}
}

// CreateClass handles creating a new class
func (s *ClassService) CreateClass(class *model.Class) (*model.Class, error) {
	// Validate UserID
	_, err := s.userRepo.FindByID(class.UserID)
	if err != nil {
		return nil, errors.New("invalid user ID for class owner")
	}

	// Validate CategoryID
	_, err = s.categoryRepo.FindByID(class.CategoryID)
	if err != nil {
		return nil, errors.New("invalid category ID")
	}

	if err := s.classRepo.Create(class); err != nil {
		return nil, errors.New("failed to create class")
	}
	return class, nil
}

// GetAllClasses retrieves all classes
func (s *ClassService) GetAllClasses() ([]model.Class, error) {
	classes, err := s.classRepo.FindAll()
	if err != nil {
		return nil, errors.New("failed to retrieve classes")
	}
	return classes, nil
}

// GetClassByID retrieves a class by ID
func (s *ClassService) GetClassByID(id uint) (*model.Class, error) {
	class, err := s.classRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return class, nil
}

// UpdateClass updates an existing class
func (s *ClassService) UpdateClass(updatedClass *model.Class, userID uint) (*model.Class, error) {
	existingClass, err := s.classRepo.FindByID(updatedClass.ID)
	if err != nil {
		return nil, errors.New("class not found")
	}

	// Only the owner can update the class
	if existingClass.UserID != userID {
		return nil, errors.New("unauthorized to update this class")
	}

	// Validate CategoryID if provided
	if updatedClass.CategoryID != 0 {
		_, err := s.categoryRepo.FindByID(updatedClass.CategoryID)
		if err != nil {
			return nil, errors.New("invalid category ID")
		}
		existingClass.CategoryID = updatedClass.CategoryID
	}

	existingClass.Title = updatedClass.Title
	existingClass.Description = updatedClass.Description

	if err := s.classRepo.Update(existingClass); err != nil {
		return nil, errors.New("failed to update class")
	}
	return existingClass, nil
}

// DeleteClass deletes a class
func (s *ClassService) DeleteClass(id uint, userID uint) error {
	class, err := s.classRepo.FindByID(id)
	if err != nil {
		return errors.New("class not found")
	}

	// Only the owner can delete the class
	if class.UserID != userID {
		return errors.New("unauthorized to delete this class")
	}

	return s.classRepo.Delete(id)
}

// EnrollInClass handles user enrollment in a class
func (s *ClassService) EnrollInClass(userID, classID uint) (*model.Enrollment, error) {
	// Check if user exists
	_, err := s.userRepo.FindByID(userID)
	if err != nil {
		return nil, errors.New("user not found")
	}

	// Check if class exists
	class, err := s.classRepo.FindByID(classID)
	if err != nil {
		return nil, errors.New("class not found")
	}

	// Prevent user from enrolling in their own class
	if class.UserID == userID {
		return nil, errors.New("cannot enroll in your own class")
	}

	// Check if already enrolled
	_, err = s.enrollmentRepo.FindByUserIDAndClassID(userID, classID)
	if err == nil {
		return nil, errors.New("already enrolled in this class")
	}

	enrollment := &model.Enrollment{
		UserID:  userID,
		ClassID: classID,
	}

	if err := s.enrollmentRepo.Create(enrollment); err != nil {
		return nil, errors.New("failed to enroll in class")
	}
	return enrollment, nil
}

// GetUserEnrollments retrieves all classes a user is enrolled in
func (s *ClassService) GetUserEnrollments(userID uint) ([]model.Enrollment, error) {
	enrollments, err := s.enrollmentRepo.FindAllByUserID(userID)
	if err != nil {
		return nil, errors.New("failed to retrieve user enrollments")
	}
	return enrollments, nil
}

// UnenrollFromClass handles unenrollment from a class
func (s *ClassService) UnenrollFromClass(userID, classID uint) error {
	// Check if enrollment exists
	enrollment, err := s.enrollmentRepo.FindByUserIDAndClassID(userID, classID)
	if err != nil {
		return errors.New("enrollment not found")
	}

	return s.enrollmentRepo.Delete(enrollment.ID)
}