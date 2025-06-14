	package models

	import "gorm.io/gorm"

	type CartItem struct {
		gorm.Model
		ProductID uint
		Quantity	int
		CartID uint

		Product Product `gorm:"foreignKey:ProductID"` //didefenisikan apabila perlu preload
	}