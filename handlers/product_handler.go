package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/seccret404/Ecomerce-Api/models"
	"github.com/seccret404/Ecomerce-Api/services"
	"github.com/seccret404/Ecomerce-Api/upload"
)

func CreateProduct(c *fiber.Ctx)error{
	name := c.FormValue("name")
	price := c.FormValue("price")
	quantity := c.FormValue("quantity")

	fileHeader, err := c.FormFile("image")
	if err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error" : "Ga ada filenya",
		})
	}

	imageUrl, err := upload.UploadFile(fileHeader)
	if err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error" : "gagal upload ke cloudinary",
		})
	}

	quantityconv, err := strconv.ParseUint(quantity, 10,32)
	if err != nil{
		return err
	}
	finalQuantity := uint(quantityconv)
	product := models.Product{
		Name: name,
		Price: price,
		Quantity: finalQuantity,
		Image: imageUrl,
	}

	newProduct, err := services.CreateProductService(&product)
	if err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error" : "ga bisa nambahin product" + err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success" : newProduct,
	})

}

func GetProduct(c *fiber.Ctx)error{
	product, err := services.GetProductService()
	if err != nil{
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error" : "Gagal dapat product",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success" : product,
	})

}

func GetByIDProject(c *fiber.Ctx)error{
	var id = c.Params("id")

	product, err := services.GetByIdProductService(id)
	if err != nil{
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error" : "Ga nemo bro",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success" : product,
	})
}

func UpdateProduct(c *fiber.Ctx)error{
	var id = c.Params("id")
	product, err := services.GetByIdProductService(id)
	if err != nil{
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error" : "tidak menemukan product itu",
		})
	}

	if name := c.FormValue("name"); name != ""{
		product.Name = name
	}
	if price := c.FormValue("price"); price != "" {
		product.Price = price
	}
	if quantity := c.FormValue("quantity"); quantity != "" {
		 
		quantity, err := strconv.ParseUint(quantity, 10, 32)
		if err != nil{
			return nil
		}
		finalQuantity := uint(quantity)
		
		product.Quantity = finalQuantity
	}

	file, err := c.FormFile("image")
	if err != nil && file !=  nil{
		imageUrl, err := upload.UploadFile(file)
		if err != nil{
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error" : "gabisa upload ke cloud",
			})
		}
		product.Image = imageUrl
	}

	updateProduct, err := services.UpdateProductService(id, product)
	if err !=  nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error" : "Gagal update product",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success" : updateProduct,
	})
}