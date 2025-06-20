package model

type Order struct {
	ID       uint   `gorm:"primaryKey"`
	UserID   uint
	ProductID uint
	Status   string
}
