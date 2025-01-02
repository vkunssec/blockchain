package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vkunssec/go-blockchain/pkg/domain"
	"github.com/vkunssec/go-blockchain/pkg/repository"

	_ "github.com/vkunssec/go-blockchain/docs"
)

type SaveBlockResponse struct {
	Message string `json:"message" example:"Block saved successfully"`
	Hash    string `json:"hash" example:"24789bede423e5c23c25856ae87bea9e37c57963ec0fbce4702a4f15cbb56a5c"`
}

type ErrorResponse struct {
	Message string `json:"message" example:"Error saving block"`
}

// @Summary Save a block
// @Description Save a block in the blockchain
// @Tags block
// @Accept json
// @Produce json
// @Param block body domain.Block true "Block"
// @Success 200 {object} SaveBlockResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router / [post]
func SaveBlock(c *fiber.Ctx) error {
	body := new(domain.Block)
	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{Message: err.Error()})
	}

	block, err := repository.SaveBlock(body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{Message: err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(SaveBlockResponse{Message: "Block saved successfully", Hash: block.Hash})
}
