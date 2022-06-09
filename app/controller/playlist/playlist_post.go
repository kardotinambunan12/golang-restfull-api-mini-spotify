package playlist

import (
	"final-project/app/model"
	"final-project/app/services"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// PostPlaylist godoc
// @Summary Create new Playlist
// @Description Create new Playlist
// @Accept  application/json
// @Produce  application/json
// @Success 201 {object} model.PlayListApi "Created"
// @Success 400
// @Success 401
// @Param data body model.PlayListApi true "Playlist data"
// @Router /playlist [post]
// @Tags Playlist
func PlaylistPost(c *fiber.Ctx) error {
	db := services.DB
	var playlist model.PlayListApi

	accept := c.Get("accept")

	if !strings.EqualFold(accept, "application/json") {
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid Heades",
		})
	}

	if err := c.BodyParser(&playlist); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid Request",
		})
	}
	create_playlist := &model.PlayList{PlayListApi: playlist}
	err := db.Model(&model.PlayList{}).Preload("Gendre").Preload("Artist").Create(create_playlist)
	if err != nil {
		c.Status(400).JSON(fiber.Map{
			"message": err,
			"data":    nil,
		})
	}
	return c.Status(201).JSON(fiber.Map{
		"message": "success",
		"status":  200,
		"items":   create_playlist,
	})
}
