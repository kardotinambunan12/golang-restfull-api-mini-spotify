package album

import (
	"final-project/app/model"
	"final-project/app/services"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// PostAlbum godoc
// @Summary Create new Album
// @Description Create new Album
// @Accept  application/json
// @Produce  application/json
// @Success 201 {object} model.AlbumApi "Created"
// @Success 400
// @Success 401
// @Param data body model.AlbumApi true "Album data"
// @Router /album [post]
// @Tags Album
func AlbumPost(c *fiber.Ctx) error {
	db := services.DB
	var album model.AlbumApi

	accept := c.Get("accept")

	if !strings.EqualFold(accept, "application/json") {
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid Heades",
		})
	}

	if err := c.BodyParser(&album); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid Request",
		})
	}
	create_album := &model.Album{AlbumApi: album}
	err := db.Model(&model.Album{}).Preload("Gendre").Preload("Artist").Create(create_album)
	if err != nil {
		c.Status(400).JSON(fiber.Map{
			"message": err,
			"data":    nil,
		})
	}
	return c.Status(201).JSON(fiber.Map{
		"message": "success",
		"status":  200,
		"items":   create_album,
	})
}
