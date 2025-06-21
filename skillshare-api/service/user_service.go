package service

import (
	"errors"
	"skillshare-api/helper"
	"skillshare-api/model"
	"skillshare-api/repository"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

// RegisterUser handles the business logic for user registration
func (s *UserService) RegisterUser(user *model.User) (*model.User, error) {
	// Hash password
	hashedPassword, err := helper.HashPassword(user.Password)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}
	user.Password = hashedPassword

	// Save user to database
	if err := s.userRepo.Create(user); err != nil {
		return nil, errors.New("email already registered or failed to save user")
	}
	return user, nil
}

// LoginUser handles the business logic for user login
func (s *UserService) LoginUser(email, password string) (string, error) {
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	if !helper.CheckPasswordHash(password, user.Password) {
		return "", errors.New("invalid credentials")
	}

	token, err := helper.GenerateJWT(user.ID, user.Email)
	if err != nil {
		return "", errors.New("failed to generate token")
	}
	return token, nil
}

// GetUserByID retrieves a user by ID
func (s *UserService) GetUserByID(id uint) (*model.User, error) {
	user, err := s.userRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// UpdateUser updates an existing user
func (s *UserService) UpdateUser(user *model.User) (*model.User, error) {
	existingUser, err := s.userRepo.FindByID(user.ID)
	if err != nil {
		return nil, errors.New("user not found")
	}

	// Update fields
	existingUser.Name = user.Name
	existingUser.Email = user.Email
	// Only update password if provided
	if user.Password != "" {
		hashedPassword, err := helper.HashPassword(user.Password)
		if err != nil {
			return nil, errors.New("failed to hash new password")
		}
		existingUser.Password = hashedPassword
	}

	if err := s.userRepo.Update(existingUser); err != nil {
		return nil, errors.New("failed to update user")
	}
	return existingUser, nil
}

// DeleteUser deletes a user
func (s *UserService) DeleteUser(id uint) error {
	return s.userRepo.Delete(id)
}