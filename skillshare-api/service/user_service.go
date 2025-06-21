package service

import (
	"errors"
	"skillshare-api/model"
	"skillshare-api/repository"
	// Hapus import "golang.org/x/crypto/bcrypt" jika tidak digunakan lagi (pastikan ini sudah dihapus jika tidak dipakai)
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

// RegisterUser handles user registration.
// IMPORTANT: This version stores passwords in plaintext, which is HIGHLY INSECURE.
// It's used here only for demonstration as per your request.
func (s *UserService) RegisterUser(user *model.User) (*model.User, error) {
	// Check if email already exists
	existingUser, err := s.userRepo.FindByEmail(user.Email)

	// If a user with this email was found, return an "already registered" error.
	if existingUser != nil {
		return nil, errors.New("email already registered")
	}

	// If there was an error from FindByEmail that is NOT "user not found",
	// it means a real database error occurred, so we return that error.
	// We explicitly check for "user not found" string to handle cases where
	// errors.Is might not perfectly match due to underlying error types.
	if err != nil && err.Error() != "user not found" {
		return nil, err
	}

	// At this point:
	// - either existingUser is nil and err is "user not found" (meaning email is available)
	// - or existingUser is nil and err is nil (meaning email is available and repo didn't return an error)

	// Store password as plaintext (EXTREMELY INSECURE - FOR DEMO ONLY)
	// No hashing logic is applied here.
	if err := s.userRepo.Create(user); err != nil {
		return nil, errors.New("failed to register user")
	}
	return user, nil
}

// LoginUser handles user login by comparing plaintext passwords.
// IMPORTANT: This version compares passwords in plaintext, which is HIGHLY INSECURE.
// It's used here only for demonstration as per your request.
func (s *UserService) LoginUser(email, password string) (*model.User, error) {
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		// If user not found or any other error from repository, treat as invalid credentials
		return nil, errors.New("invalid credentials")
	}

	// Compare plaintext password directly (EXTREMELY INSECURE - FOR DEMO ONLY)
	if user.Password != password {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}

// GetUserByID retrieves a user by ID.
func (s *UserService) GetUserByID(id uint) (*model.User, error) {
	user, err := s.userRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// UpdateUser updates an existing user.
func (s *UserService) UpdateUser(updatedUser *model.User) (*model.User, error) {
	existingUser, err := s.userRepo.FindByID(updatedUser.ID)
	if err != nil {
		return nil, errors.New("user not found")
	}

	existingUser.Name = updatedUser.Name
	existingUser.Email = updatedUser.Email
	// If you want to allow password updates without hashing, uncomment the block below.
	// Remember this is insecure in plaintext.
	// if updatedUser.Password != "" {
	// 	existingUser.Password = updatedUser.Password
	// }

	if err := s.userRepo.Update(existingUser); err != nil {
		return nil, errors.New("failed to update user")
	}
	return existingUser, nil
}

// DeleteUser deletes a user by ID.
func (s *UserService) DeleteUser(id uint) error {
	return s.userRepo.Delete(id)
}