package song

import (
	"final-project/app/model"
	"final-project/app/services"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// SongtGet godoc
// @Summary Get List of Song
// @Description Get List of Song
// @Accept application/json
// @Produce application/json
// @Success 200 {object} []model.Song "Success"
// @Router /song [get]
// @Tags Song
func SongGet(c *fiber.Ctx) error {
	db := services.DB

	accept := c.Get("accept")
	if !strings.EqualFold(accept, "application/json") {
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid accept header",
		})
	}
	var song []model.Song
	db.Model(&model.Song{}).
		Preload("Artist").
		Find(&song)

	return c.Status(200).JSON(fiber.Map{
		"message": "success",
		"status":  200,
		"items":   song,
	})
}
