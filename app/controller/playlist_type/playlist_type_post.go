package playlist_type

import (
	"final-project/app/model"
	"final-project/app/services"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// @PostPlaylistType go doc
// @Summary Create new PlaylistType
// @Description create new PlaylistType
// @Accept application/json
// @Produce application/json
// @Success 201 {object} model.PlaylistTypeApi "created"
// @Success 400
// @Success 401
// @Param data body model.PlaylistTypeApi true "PlaylistType data"
// @Router /playlist_type [post]
// @Tags PlaylistType
func PlaylistTypePost(c *fiber.Ctx) error {
	db := services.DB

	accept := c.Get("accept")
	if !strings.EqualFold(accept, "application/json") {
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid header",
		})
	}
	var playlist_type model.PlaylistTypeApi

	if err := c.BodyParser(&playlist_type); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid request",
		})
	}
	create_playlist_type := &model.PlaylistType{PlaylistTypeApi: playlist_type}
	db.Model(&model.PlaylistType{}).Create(create_playlist_type)

	return c.Status(201).JSON(create_playlist_type)

}
