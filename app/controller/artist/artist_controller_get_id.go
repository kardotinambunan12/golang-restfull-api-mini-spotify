package artist

import (
	"final-project/app/model"
	"final-project/app/services"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// ArtistGet godoc
// @Summary Get Artist by id
// @Description Get Artist by id
// @Accept application/json
// @Produce application/json
// @Success 200 {object} []model.Artist "Success"
// @Success 400
// @Success 404
// @Param id path string true "Artist ID"
// @Router /artist/{id} [get]
// @Tags Artist
func ArtistGetById(c *fiber.Ctx) error {
	db := services.DB
	var artist model.Artist

	accept := c.Get("accept")
	if !strings.EqualFold(accept, "application/json") {
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid accept header",
		})
	}
	id := c.Params("id")
	db.Where(`id=?`, id).First(&artist)

	return c.Status(200).JSON(artist)
}
