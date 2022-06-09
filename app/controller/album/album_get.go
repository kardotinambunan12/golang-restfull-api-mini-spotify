package album

import (
	"final-project/app/model"
	"final-project/app/services"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// AlbumGet godoc
// @Summary Get List of Album
// @Description Get List of Album
// @Accept application/json
// @Produce application/json
// @Success 200 {object} []model.Album "Success"
// @Router /album [get]
// @Tags Album
func AlbumGetAll(c *fiber.Ctx) error {
	db := services.DB
	var album []model.Album
	accept := c.Get("accept")
	if !strings.EqualFold(accept, "application/json") {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid header",
		})
	}
	db.Model(&model.Album{}).
		Preload("Gendre").
		Preload("Artist").
		Find(&album)
	return c.Status(200).JSON(fiber.Map{
		"message": "success",
		"status":  200,
		"items":   album,
	})
}
