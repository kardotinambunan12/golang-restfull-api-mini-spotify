package playlist_type

import (
	"final-project/app/model"
	"final-project/app/services"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// PutPlaylistType Go doc
// @Summary Updated PlaylistType by Id
// @Description udpated PlaylistType by Id
// @Accept application/json
// @Produce application/json
// @Success 201 {object} model.PlaylistTypeApi "OK"
// @Success 401
// @Success 404
// @Param id path string true "PlaylistType ID"
// @Param data body model.PlaylistTypeApi true "PlaylistType data"
// @Router /playlist_type/{id} [put]
// @Tags PlaylistType
func PlaylistTypePut(c *fiber.Ctx) error {

	db := services.DB

	accept := c.Get("accept")

	if !strings.EqualFold(accept, "application/json") {
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid headers",
		})
	}
	var playlist_type model.PlaylistTypeApi
	id := c.Params("id")

	if err := c.BodyParser(&playlist_type); err != nil {
		fmt.Println("error :", err.Error())
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid request",
		})
	}

	data_update := &model.PlaylistType{PlaylistTypeApi: playlist_type}
	result := db.Model(&model.PlaylistType{}).Where(`id=?`, id).Updates(data_update)
	data_row := result.RowsAffected
	msg := fmt.Sprintf(`%d row affected`, data_row)

	return c.JSON(fiber.Map{
		"message": msg,
		"status":  201,
		"data":    data_update,
	})
}
