package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	SessionID string
	CartItems []CartItem `gorm:"foreignKey:CartID"`
	IsCheckOut bool
}