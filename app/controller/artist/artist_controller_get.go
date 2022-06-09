package artist

import (
	"final-project/app/model"
	"final-project/app/services"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// ArtistGet godoc
// @Summary Get List of Artist
// @Description Get List of Artist
// @Accept application/json
// @Produce application/json
// @Success 200 {object} []model.Artist "Success"
// @Router /artist [get]
// @Tags Artist
func ArtistGet(c *fiber.Ctx) error {
	db := services.DB
	var artist []model.Artist

	accept := c.Get("accept")
	if !strings.EqualFold(accept, "application/json") {
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid accept header",
		})
	}
	db.Model(&model.Artist{}).Find(&artist)
	// if err != nil {
	// 	return c.Status(500).JSON(fiber.Map{
	// 		"message": err,
	// 		"data":    nil,
	// 	})
	// }

	return c.Status(200).JSON(fiber.Map{
		"message": "success",
		"status":  200,
		"items":   artist,
	})
}
