package playlist_type

import (
	"final-project/app/model"
	"final-project/app/services"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// @PlaylistTypeGetById go doc
// @Summary get PlaylistType by ID
// @Description get PlaylistType by ID
// @Accept application/json
// @Produce application/json
// @Success 200 {object} model.PlaylistType "Success"
// @Success 400
// @Success 404
// @Param id path string true  "PlaylistType ID"
// @Router /playlist_type/{id} [get]
// @Tags PlaylistType
func PlaylistTypeGetById(c *fiber.Ctx) error {
	db := services.DB
	var playlist_type model.PlaylistType

	accept := c.Get("accept")
	if !strings.EqualFold(accept, "application/json") {
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid header",
		})
	}
	id := c.Params("id")

	db.Where("id=?", id).First(&playlist_type)

	return c.JSON(fiber.Map{
		"message": "sucess",
		"items":   playlist_type,
	})
}
