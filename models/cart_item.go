package models

import "gorm.io/gorm"

type CartItem struct {
	gorm.Model
	ProductID uint
	Quantity	int
	Product	Product `gorm:"foreignKey:ProductID"`
}