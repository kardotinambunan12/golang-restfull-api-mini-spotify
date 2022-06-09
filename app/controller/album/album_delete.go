package album

import (
	"final-project/app/model"
	"final-project/app/services"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

//Deleted Album go Doc
// @Summary Delete Album by id
// @Description Delete Album by id
// @Accept application/json
// @Produce application/json
// @Success 200
// @Success 400
// @Success 404
// @Param id path string true "Album ID"
// @Router /album/{id} [delete]
// @Tags Album
func AlbumDelete(c *fiber.Ctx) error {
	db := services.DB

	accept := c.Get("accept")
	if !strings.EqualFold(accept, "application/json") {
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid accept header",
		})
	}
	id := c.Params("id")
	response := db.Where(`id=?`, id).Delete(&model.Album{})
	data := response.RowsAffected
	msg := fmt.Sprintf(`album  %d has deleted`, data)

	return c.Status(200).JSON(fiber.Map{
		"status":  200,
		"message": msg,
	})
}
