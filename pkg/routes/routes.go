package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vkunssec/blockchain/pkg/handlers"

	_ "github.com/vkunssec/blockchain/docs"
)

// @Summary Salva um bloco
// @Description Salva um bloco na blockchain
func SetupRoutes(app *fiber.App) {
	app.Post("/", handlers.SaveBlock)
}
