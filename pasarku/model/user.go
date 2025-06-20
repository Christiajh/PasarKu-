package model

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"not null"`
	Email    string `gorm:"unique"`
	Password string
	Role     string `gorm:"default:user"`
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
