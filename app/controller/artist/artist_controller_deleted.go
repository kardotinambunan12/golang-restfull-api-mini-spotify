package artist

import (
	"final-project/app/model"
	"final-project/app/services"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

//Deleted Artist go Doc
// @Summary Delete Artits by id
// @Description Delete Artits by id
// @Accept application/json
// @Produce application/json
// @Success 200
// @Success 400
// @Success 404
// @Param id path string true "Artits ID"
// @Router /artist/{id} [delete]
// @Tags Artist
func ArtistDeleted(c *fiber.Ctx) error {
	db := services.DB

	content_type := c.Get("accept")
	if !strings.EqualFold(content_type, "application/json") {
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid accept header",
		})
	}

	id := c.Params("id")
	response := db.Where(`id=?`, id).Delete(&model.Artist{})
	data := response.RowsAffected

	msg := fmt.Sprintf(`artist  %d has deleted`, data)

	return c.Status(200).JSON(fiber.Map{
		"message": msg,
	})
}
