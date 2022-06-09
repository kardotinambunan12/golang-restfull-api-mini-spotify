package gendre

import (
	"final-project/app/model"
	"final-project/app/services"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// GendreGet godoc
// @Summary Get List of Gendre
// @Description Get List of Gendre
// @Accept application/json
// @Produce application/json
// @Success 200 {object} []model.Gendre "Success"
// @Router /gendre [get]
// @Tags Gendre
func GendreGetAll(c *fiber.Ctx) error {

	db := services.DB
	var gendre []model.Gendre
	accept := c.Get("accept")
	if !strings.EqualFold(accept, "application/json") {
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid headers",
		})
	}
	db.Model(&model.Gendre{}).Find(&gendre)

	return c.Status(200).JSON(fiber.Map{
		"message": "success",
		"status":  200,
		"items":   gendre,
	})
}
