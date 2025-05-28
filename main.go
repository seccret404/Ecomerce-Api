package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/seccret404/Ecomerce-Api/config"
	"github.com/seccret404/Ecomerce-Api/routes"
	// "github.com/seccret404/Ecomerce-Api/models"
)

func main() {
	fmt.Println("Hello World")

	//load env
	err := godotenv.Load()
	if err != nil{
		log.Fatal("gagal buka env")
	}

	//database
	config.InitDatabase()
	config.InitRedis()

	// models.MigateALL()
	app := fiber.New()

	routes.RoutesSetUp(app)

	log.Fatal(app.Listen(":3000"))
}