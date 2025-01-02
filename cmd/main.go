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
// @description API Para uso de Blockchain
// @host localhost:3000
// @BasePath /
func main() {
	// blockchain := repository.Blockchain{
	// 	Difficulty: 3,
	// }

	// block := *repository.NewBlock("First Node", []*repository.Block{
	// 	{
	// 		Hash: "24789bede423e5c23c25856ae87bea9e37c57963ec0fbce4702a4f15cbb56a5c",
	// 	},
	// })

	// blockchain.AddBlock(&block)

	// block2 := *repository.NewBlock("Second Node", blockchain.Chain)

	// blockchain.AddBlock(&block2)

	// r, _ := json.MarshalIndent(&blockchain, "", "  ")
	// log.Println(string(r))

	app := fiber.New()

	middleware.SetupCommonMiddleware(app)
	middleware.SetupSwagger(app)
	routes.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
