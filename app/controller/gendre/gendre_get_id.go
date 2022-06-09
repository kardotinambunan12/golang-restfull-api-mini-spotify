package gendre

import (
	"final-project/app/model"
	"final-project/app/services"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// @GendreGetById go doc
// @Summary get gendre by ID
// @Description get gendre by ID
// @Accept application/json
// @Produce application/json
// @Success 200 {object} model.Gendre "Success"
// @Success 400
// @Success 404
// @Param id path string true  "Gendre ID"
// @Router /gendre/{id} [get]
// @Tags Gendre
func GendreGetById(c *fiber.Ctx) error {

	db := services.DB
	var gendre model.Gendre

	accept := c.Get("accept")
	if !strings.EqualFold(accept, "application/json") {
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid header",
		})
	}
	id := c.Params("id")

	db.Where("id=?", id).First(&gendre)

	return c.JSON(fiber.Map{
		"message": "sucess",
		"items":   gendre,
	})

}
