package playlist

import (
	"final-project/app/model"
	"final-project/app/services"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// @PlaylistGetById go doc
// @Summary get Playlist by ID
// @Description get Playlist by ID
// @Accept application/json
// @Produce application/json
// @Success 200 {object} model.PlayList "Success"
// @Success 400
// @Success 404
// @Param id path string true  "PlayList ID"
// @Router /playlist/{id} [get]
// @Tags Playlist
func PlaylistGetById(c *fiber.Ctx) error {
	db := services.DB

	accept := c.Get("accept")
	if !strings.EqualFold(accept, "application/json") {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid headers",
		})
	}
	var playlist model.PlayList

	id := c.Params("id")
	db.Where(`id=?`, id).
		Preload("PlaylistType").
		First(&playlist)
	return c.JSON(fiber.Map{
		"message": "sucess",
		"items":   playlist,
	})
}
