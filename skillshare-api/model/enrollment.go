package model

import "gorm.io/gorm"

// Enrollment represents a user's enrollment in a class
type Enrollment struct {
	gorm.Model
	UserID  uint `json:"user_id"`
	User    User `json:"-" gorm:"foreignKey:UserID"` // sembunyikan agar tidak rekursif

	ClassID uint  `json:"class_id"`
	Class   Class `json:"-" gorm:"foreignKey:ClassID"` // sembunyikan juga
}
