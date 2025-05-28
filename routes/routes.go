package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/seccret404/Ecomerce-Api/handlers"
)

func RoutesSetUp(app *fiber.App) {
	routeGroup := app.Group("/api")

	routeGroup.Post("/product", handlers.CreateProduct)
	routeGroup.Get("/product", handlers.GetProduct)
}