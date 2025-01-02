package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vkunssec/go-blockchain/pkg/handlers"

	_ "github.com/vkunssec/go-blockchain/docs"
)

// @Summary Save a block
// @Description Save a block in the blockchain
func SetupRoutes(app *fiber.App) {
	app.Post("/", handlers.SaveBlock)
}
