package playlist_type

import (
	"final-project/app/model"
	"final-project/app/services"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// PlaylistTypeGet godoc
// @Summary Get List of PlaylistType
// @Description Get List of PlaylistType
// @Accept application/json
// @Produce application/json
// @Success 200 {object} []model.PlaylistType "Success"
// @Router /playlist_type [get]
// @Tags PlaylistType
func PlaylistTypeGet(c *fiber.Ctx) error {
	db := services.DB
	var playlist_type []model.PlaylistType
	accept := c.Get("accept")
	if !strings.EqualFold(accept, "application/json") {
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid headers",
		})
	}
	db.Model(&model.PlaylistType{}).Find(&playlist_type)

	return c.JSON(fiber.Map{
		"message": "success",
		"status":  200,
		"items":   playlist_type,
	})
}
