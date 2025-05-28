package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name		string 
	Price	string
	Quantity	uint
	Image	string
	CartItems	[]CartItem `gorm:"foreginKey:ProductID"`

}