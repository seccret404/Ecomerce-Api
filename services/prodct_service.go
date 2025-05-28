package services

import (
	"context"
	"encoding/json"
	"time"

	"github.com/seccret404/Ecomerce-Api/config"
	"github.com/seccret404/Ecomerce-Api/models"
)

func CreateProductService(product *models.Product)(*models.Product, error){
	if err := config.DB.Create(product).Error; err != nil{
		return nil, err
	}

	return product, nil
}

func GetProductService()([]models.Product, error){
	ctx := context.Background()
	var product []models.Product
	//kalau mau buat get by id... buat var baru buat save "product" + id as a key
	// cek redis
	cached, err := config.RedisClient.Get(ctx, "all_product").Result()
	if err == nil{
		//jika ada data di redis akan di alihkan ke product
		if err := json.Unmarshal([]byte(cached), &product); err != nil{
			return product, nil
		}
	}

	//semisal ga ada di redis maka akan kesini
	if err := config.DB.Find(&product).Error; err !=nil{
		return nil, err
	}

	//data akan disimpan ke redis setelah di ambil dari database
	jsonData, err := json.Marshal(product)
	if err == nil {
		config.RedisClient.Set(ctx, "all_product", jsonData, time.Minute*1)
	}

	return product, nil

}

func GetByIdProductService(id string)(*models.Product, error){
	var product models.Product
	if err := config.DB.First(&product, id).Error; err != nil{
		return nil, err
	}
	return &product, nil
}

func UpdateProductService(id string, updateProduct *models.Product)(*models.Product, error){
	var product models.Product
	if err := config.DB.First(&product, id).Error; err != nil{
		return nil, err
	}

	product.Name = updateProduct.Name
	product.Price = updateProduct.Price
	product.Quantity = updateProduct.Quantity

	if updateProduct.Image != ""{
		product.Image = updateProduct.Image
	}

	if err := config.DB.Save(&product).Error; err != nil{
		return nil, err
	}

	return &product, nil
}