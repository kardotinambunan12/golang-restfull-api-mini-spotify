package song

import (
	"final-project/app/model"
	"final-project/app/services"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// PostSong godoc
// @Summary Create new Song
// @Description Create new Song
// @Accept  application/json
// @Produce  application/json
// @Success 201 {object} model.SongApi "Created"
// @Success 400
// @Success 401
// @Param data body model.SongApi true "Song data"
// @Router /song [post]
// @Tags Song
func SongPost(c *fiber.Ctx) error {
	db := services.DB
	var song model.SongApi

	accept := c.Get("accept")
	if !strings.EqualFold(accept, "application/json") {
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid accept header",
		})
	}
	if err := c.BodyParser(&song); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid Request",
		})
	}

	create_song := &model.Song{SongApi: song}
	err := db.Model(&model.Song{}).
		Preload("Artist").
		Create(create_song)
	if err != nil {
		c.Status(400).JSON(fiber.Map{
			"message": err,
			"data":    nil,
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "success",
		"status":  200,
		"data":    create_song,
	})
}
