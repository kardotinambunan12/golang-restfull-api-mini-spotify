package album

import (
	"final-project/app/model"
	"final-project/app/services"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// @AlbumGetById go doc
// @Summary get Album by ID
// @Description get Album by ID
// @Accept application/json
// @Produce application/json
// @Success 200 {object} model.Album "Success"
// @Success 400
// @Success 404
// @Param id path string true  "Album ID"
// @Router /album/{id} [get]
// @Tags Album
func AlbumGetById(c *fiber.Ctx) error {

	db := services.DB

	accept := c.Get("accept")
	if !strings.EqualFold(accept, "application/json") {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid headers",
		})
	}
	var album model.Album

	id := c.Params("id")
	db.Where(`id=?`, id).
		Preload("Gendre").
		Preload("Artist").
		First(&album)
	return c.JSON(fiber.Map{
		"message": "sucess",
		"items":   album,
	})
}
