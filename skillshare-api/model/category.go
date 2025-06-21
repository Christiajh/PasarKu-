package model

import "gorm.io/gorm"

// Category represents a category for classes
type Category struct {
	gorm.Model
	Name    string  `json:"name" gorm:"unique"`
	Classes []Class `json:"classes" gorm:"foreignKey:CategoryID"`
}