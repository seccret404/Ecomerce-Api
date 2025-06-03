package handlers

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/seccret404/Ecomerce-Api/helper"
	"github.com/seccret404/Ecomerce-Api/models"
	"github.com/seccret404/Ecomerce-Api/services"
)

func GetCart(c *fiber.Ctx)error{
	session_id := helper.GetSessionID(c)

	cart, err := services.GetCartService(session_id)
	if err != nil{
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error" : "Ga ketemu",
		})
	}

	if len(cart.CartItems) == 0{
		return c.JSON(fiber.Map{
			"message" : "new cart udah dibuat",
			"cart" : cart,
		})
	}

	return c.JSON(cart)
	
}

func AddCart(c *fiber.Ctx)error{
	seesion_id := helper.GetSessionID(c)

	var input struct{
		ProductID uint `json:"product_id"`
		Quantity int `json:"quantity"`
	}

	if err:= c.BodyParser(&input); err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error" : "request tidak valid",
		})
	}

	//validasi
	if input.ProductID == 0 || input.Quantity <= 0{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error" : "product id dan quantity harus valid",
		})
	}

	//cart item dari inpnut
	cartItem := models.CartItem{
		ProductID: input.ProductID,
		Quantity: input.Quantity,
	}

	item, err := services.AddCartSrvice(seesion_id, cartItem)
	if err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error" : "Gagal nambahkan data cok",
		})
	}

	return c.JSON(fiber.Map{
		"message" : "berhasil add to cart",
		"item" : item,
	})

}

func GetCartBySession(c *fiber.Ctx)error{
	session_id := helper.GetSessionID(c)
	fmt.Println("SESSION_ID:", session_id)

	items, err := services.GetCartBySessionService(session_id)
	if err != nil{
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error" : "Tidak menemukan cart",	
		})
	}

	if len(items) == 0{
		return c.JSON(fiber.Map{
			"message" : "cart kosong",
			"items" : items,
		})
	}

	return c.JSON(items)

}

func UpdateCart(c *fiber.Ctx)error{
	itemID, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil{
		return c.JSON(fiber.Map{
			"error" : "ID gak valid bro",
		})
	}

	var body struct {
		Quantity int `json:"quantity"`
	}

	if err:= c.BodyParser(&body); err !=  nil || body.Quantity <= 0{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "quantity tidak boleh 0",
		})
	}
	
	//validasi aman? lanjut ke simpan perubahan -> call service
	item, err := services.UpdateCartService(uint(itemID), body.Quantity)
	if err !=  nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error" : "Gagal nyimpan data",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message" : "Data Terupdate",
		"items" : item,
	})
}