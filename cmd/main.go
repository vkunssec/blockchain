package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/vkunssec/go-blockchain/pkg/middleware"
	"github.com/vkunssec/go-blockchain/pkg/routes"

	_ "github.com/vkunssec/go-blockchain/docs"
)

// @title Blockchain API
// @version 1.0
// @description Blockchain API
// @host localhost:3000
// @BasePath /
func main() {
	app := fiber.New()

	middleware.SetupCommonMiddleware(app)
	middleware.SetupSwagger(app)
	routes.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
