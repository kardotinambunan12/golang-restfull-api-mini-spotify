package playlist_type

import (
	"final-project/app/model"
	"final-project/app/services"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

//Deleted PlaylistType go Doc
// @Summary Delete PlaylistType by id
// @Description Delete PlaylistType by id
// @Accept application/json
// @Produce application/json
// @Success 200
// @Success 400
// @Success 404
// @Param id path string true "PlaylistType ID"
// @Router /playlist_type/{id} [delete]
// @Tags PlaylistType
func PlaylistTypeDelete(c *fiber.Ctx) error {
	db := services.DB

	accept := c.Get("accept")
	if !strings.EqualFold(accept, "application/json") {
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid accept header",
		})
	}
	id := c.Params("id")
	response := db.Where(`id=?`, id).Delete(&model.PlaylistType{})
	data := response.RowsAffected
	msg := fmt.Sprintf(`playlist type  %d has deleted`, data)

	return c.Status(200).JSON(fiber.Map{
		"status":  200,
		"message": msg,
	})
}
