package playlist

import (
	"final-project/app/model"
	"final-project/app/services"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// PlaylistGet godoc
// @Summary Get List of Playlist
// @Description Get List of Playlist
// @Accept application/json
// @Produce application/json
// @Success 200 {object} []model.PlayList "Success"
// @Router /playlist [get]
// @Tags Playlist
func PlaylistGetAll(c *fiber.Ctx) error {
	db := services.DB
	var playlist []model.PlayList
	accept := c.Get("accept")
	if !strings.EqualFold(accept, "application/json") {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid header",
		})
	}
	db.Model(&model.PlayList{}).
		Preload("PlaylistType").
		Find(&playlist)
	return c.Status(200).JSON(fiber.Map{
		"message": "success",
		"status":  200,
		"items":   playlist,
	})
}
