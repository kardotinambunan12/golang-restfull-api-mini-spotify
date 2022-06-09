package song

import (
	"final-project/app/model"
	"final-project/app/services"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// PutSong godoc
// @Summary Update Song by id
// @Description Update Song by id
// @Accept  application/json
// @Produce  application/json
// @Success 201 {object} model.SongApi "OK"
// @Success 401
// @Success 400
// @Param id path string true "Song ID"
// @Param data body model.SongApi true "Song data"
// @Router /song/{id} [put]
// @Tags Song
func SongPut(c *fiber.Ctx) error {
	accept := c.Get("accept")
	if !strings.EqualFold(accept, "application/json") {
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid accept header",
		})
	}

	var data model.SongApi

	db := services.DB

	id := c.Params("id")

	if err := c.BodyParser(&data); err != nil {
		fmt.Println("error :", err.Error())
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid request",
		})
	}

	update := &model.Song{SongApi: data}
	result := db.Model(&model.Song{}).Where(`id = ?`, id).Updates(update)
	rowCount := result.RowsAffected
	message := fmt.Sprintf(`%d row updated`, rowCount)

	return c.JSON(fiber.Map{
		"message": message,
		"data":    update,
	})

}
