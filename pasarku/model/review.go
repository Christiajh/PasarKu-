package model

type Review struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint
	ProductID uint
	Rating    int
	Comment   string
}
