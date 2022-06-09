package song

import (
	"final-project/app/model"
	"final-project/app/services"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// SongGetByID godoc
// @Summary Get Song by id
// @Description Get Song by id
// @Accept application/json
// @Produce application/json
// @Success 200 {object} []model.Song "Success"
// @Success 400
// @Success 404
// @Param id path string true "Song ID"
// @Router /song/{id} [get]
// @Tags Song
func SongGetById(c *fiber.Ctx) error {
	db := services.DB
	var song model.Song

	accept := c.Get("accept")
	if !strings.EqualFold(accept, "application/json") {
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid accept header",
		})
	}

	id := c.Params("id")
	db.Where(`id=?`, id).
		Preload("Artist").
		First(&song)

	return c.Status(200).JSON(song)
}
