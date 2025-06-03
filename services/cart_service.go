package services

import (
	"github.com/seccret404/Ecomerce-Api/config"
	"github.com/seccret404/Ecomerce-Api/models"
	"gorm.io/gorm"
)

func GetCartService(sessionID string) (*models.Cart, error) {
	var cart models.Cart
	err := config.DB.Where("session_id = ? AND is_check_out = false", sessionID).First(&cart).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			//create new
			newCart := models.Cart{
				SessionID:  sessionID,
				IsCheckOut: false,
			}
			if err := config.DB.Create(&newCart).Error; err != nil {
				return nil, err
			}
			return &newCart, nil
		}
		return nil, err
	}

	return &cart, nil

}

func AddCartSrvice(sessionID string, item models.CartItem) (*models.CartItem, error) {
	cart, err := GetCartService(sessionID)
	if err != nil {
		return nil, err
	}

	//cek apakah items udah ada pada cart item
	var exCart models.CartItem
	err = config.DB.Where("cart_id = ? AND product_id = ?", cart.ID, item.ProductID).First(&exCart).Error
	if err == nil {

		exCart.Quantity += item.Quantity

		if err := config.DB.Save(&exCart).Error; err != nil {
			return nil, err
		}
		return &exCart, nil
	}

	//kalau ga ada maka create item baru
	item.CartID = cart.ID
	if err := config.DB.Create(&item).Error; err != nil {
		return nil, err
	}

	return &item, nil
}

func GetCartBySessionService(sessionID string)([]models.CartItem, error){
	var cart models.Cart
	if err := config.DB.Where("session_id = ? AND is_check_out = false", sessionID).First(&cart).Error; err != nil{
		return nil, err
	}

	var item []models.CartItem
	if err := config.DB.Preload("Product").Where("cart_id", cart.ID).Find(&item).Error; err != nil{
		return nil, err
	}

	return item,nil
}

func UpdateCartService(itemID uint, quantity int)(*models.CartItem, error){
	var item models.CartItem

	if err := config.DB.First(&item, itemID).Error; err != nil{
		return nil, err
	}

	item.Quantity = quantity //update quantity
	if err :=  config.DB.Save(&item).Error; err != nil{
		return nil, err
	}

	return &item, nil
}
