package song

import (
	"final-project/app/model"
	"final-project/app/services"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

//Deleted Song go Doc
// @Summary Delete Song by id
// @Description Delete Song by id
// @Accept application/json
// @Produce application/json
// @Success 200
// @Success 400
// @Success 404
// @Param id path string true "Song ID"
// @Router /song/{id} [delete]
// @Tags Song
func SongDelete(c *fiber.Ctx) error {
	db := services.DB

	accept := c.Get("accept")
	if !strings.EqualFold(accept, "application/json") {
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid accept header",
		})
	}
	id := c.Params("id")
	response := db.Where(`id=?`, id).Delete(&model.Song{})
	data := response.RowsAffected
	msg := fmt.Sprintf(`song  %d has deleted`, data)

	return c.Status(200).JSON(fiber.Map{
		"status":  200,
		"message": msg,
	})
}
