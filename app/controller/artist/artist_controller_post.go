package artist

import (
	"final-project/app/model"
	"final-project/app/services"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// PostArtist godoc
// @Summary Create new Artist
// @Description Create new Artist
// @Accept  application/json
// @Produce  application/json
// @Success 201 {object} model.ArtistApi "Created"
// @Success 400
// @Success 401
// @Param data body model.ArtistApi true "Artist data"
// @Router /artist [post]
// @Tags Artist
func ArtistPost(c *fiber.Ctx) error {
	db := services.DB

	content_type := c.Get("Content-Type")
	if !strings.EqualFold(content_type, "application/json") {
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid Content-type header",
		})
	}
	var data model.ArtistApi
	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}

	create_new := &model.Artist{ArtistApi: data}
	db.Model(&model.Artist{}).Create(create_new)

	return c.Status(201).JSON(create_new)

}
