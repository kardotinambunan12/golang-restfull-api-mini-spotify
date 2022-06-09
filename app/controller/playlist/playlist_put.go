package playlist

import (
	"final-project/app/model"
	"final-project/app/services"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// PutPlaylist godoc
// @Summary Update Playlist by id
// @Description Update Playlist by id
// @Accept  application/json
// @Produce  application/json
// @Success 201 {object} model.PlayListApi "OK"
// @Success 401
// @Success 400
// @Param id path string true "Playlist ID"
// @Param data body model.PlayListApi true "Playlist data"
// @Router /playlist/{id} [put]
// @Tags Playlist
func PlaylistPut(c *fiber.Ctx) error {
	db := services.DB
	var playlist model.PlayListApi

	accept := c.Get("accept")
	if !strings.EqualFold(accept, "application/json") {
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid accept header",
		})
	}

	id := c.Params("id")

	if err := c.BodyParser(&playlist); err != nil {
		fmt.Println("error :", err.Error())
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid Request",
		})
	}

	updates := &model.PlayList{PlayListApi: playlist}
	result := db.Model(&model.PlayList{}).Where(`id=?`, id).Updates(updates)
	row_value := result.RowsAffected
	message := fmt.Sprintf(`%d row updated`, row_value)

	return c.JSON(fiber.Map{
		"message": message,
		"data":    updates,
	})
}
