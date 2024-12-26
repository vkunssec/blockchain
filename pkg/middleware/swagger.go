package middleware

import (
	"github.com/MarceloPetrucio/go-scalar-api-reference"
	"github.com/gofiber/fiber/v2"
	"github.com/vkunssec/blockchain/docs"
)

func SetupSwagger(app *fiber.App) {
	swaggerJSON := docs.SwaggerInfo.ReadDoc()

	app.Get("/docs/swagger.json", func(c *fiber.Ctx) error {
		return c.SendString(swaggerJSON)
	})

	app.Get("/swagger/*", func(ctx *fiber.Ctx) error {
		host := ctx.BaseURL()

		html, err := scalar.ApiReferenceHTML(&scalar.Options{
			SpecURL: host + "/docs/swagger.json",
			CustomOptions: scalar.CustomOptions{
				PageTitle: "API Documentation",
			},
			DarkMode: true,
		})
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		ctx.Set("Content-Type", "text/html; charset=utf-8")
		return ctx.SendString(html)
	})
}
