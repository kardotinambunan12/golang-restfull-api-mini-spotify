package gendre

import (
	"final-project/app/model"
	"final-project/app/services"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// @PostGendre go doc
// @Summary Create new Gendre
// @Description create new gendre
// @Accept application/json
// @Produce application/json
// @Success 201 {object} model.GendreApi "created"
// @Success 400
// @Success 401
// @Param data body model.GendreApi true "Gendre data"
// @Router /gendre [post]
// @Tags Gendre
func GendrePost(c *fiber.Ctx) error {

	db := services.DB

	accept := c.Get("accept")
	if !strings.EqualFold(accept, "application/json") {
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid header",
		})
	}
	var gendre model.GendreApi

	if err := c.BodyParser(&gendre); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid request",
		})
	}
	create_gendre := &model.Gendre{GendreApi: gendre}
	db.Model(&model.Gendre{}).Create(create_gendre)

	return c.Status(201).JSON(create_gendre)

}
