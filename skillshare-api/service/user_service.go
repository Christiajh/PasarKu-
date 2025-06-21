package service

import (
	"errors"
	"skillshare-api/model"
	"skillshare-api/repository"
	// Hapus import "golang.org/x/crypto/bcrypt" jika tidak digunakan lagi
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

// RegisterUser handles user registration without password hashing
func (s *UserService) RegisterUser(user *model.User) (*model.User, error) {
	// Check if email already exists
	existingUser, err := s.userRepo.FindByEmail(user.Email)
	if existingUser != nil {
		return nil, errors.New("email already registered")
	}
	if err != nil && !errors.Is(err, errors.New("user not found")) { // Handle actual errors, not just "not found"
		return nil, err
	}

	// Store password as plaintext (EXTREMELY INSECURE - FOR DEMO ONLY)
	// Tidak perlu hashing lagi
	// user.Password = string(hashedPassword) // Baris ini dihapus

	if err := s.userRepo.Create(user); err != nil {
		return nil, errors.New("failed to register user")
	}
	return user, nil
}

// LoginUser handles user login without password hashing verification
func (s *UserService) LoginUser(email, password string) (*model.User, error) {
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		if errors.Is(err, errors.New("user not found")) {
			return nil, errors.New("invalid credentials")
		}
		return nil, err
	}

	// Compare plaintext password directly (EXTREMELY INSECURE - FOR DEMO ONLY)
	if user.Password != password { // Langsung bandingkan password
		return nil, errors.New("invalid credentials")
	}

	return user, nil
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
func (s *UserService) UpdateUser(updatedUser *model.User) (*model.User, error) {
	existingUser, err := s.userRepo.FindByID(updatedUser.ID)
	if err != nil {
		return nil, errors.New("user not found")
	}

	existingUser.Name = updatedUser.Name
	existingUser.Email = updatedUser.Email
	// Jika ada opsi untuk mengganti password tanpa hash:
	// if updatedUser.Password != "" {
	// 	existingUser.Password = updatedUser.Password
	// }

	if err := s.userRepo.Update(existingUser); err != nil {
		return nil, errors.New("failed to update user")
	}
	return existingUser, nil
}

// DeleteUser deletes a user
func (s *UserService) DeleteUser(id uint) error {
	return s.userRepo.Delete(id)
}