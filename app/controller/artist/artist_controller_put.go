package artist

import (
	"final-project/app/model"
	"final-project/app/services"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// PutArtist godoc
// @Summary Update Artist by id
// @Description Update Artist by id
// @Accept  application/json
// @Produce  application/json
// @Success 201 {object} model.ArtistApi "OK"
// @Success 401
// @Success 400
// @Param id path string true "Artist ID"
// @Param data body model.ArtistApi true "Artist data"
// @Router /artist/{id} [put]
// @Tags Artist
func ArtistPut(c *fiber.Ctx) error {

	accept := c.Get("accept")
	if !strings.EqualFold(accept, "application/json") {
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid accept header",
		})
	}

	var artist model.ArtistApi

	db := services.DB

	id := c.Params("id")

	if err := c.BodyParser(&artist); err != nil {
		fmt.Println("error :", err.Error())
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid request",
		})
	}

	data_update := &model.Artist{ArtistApi: artist}
	result := db.Model(&model.Artist{}).Where(`id = ?`, id).Updates(data_update)

	data_row := result.RowsAffected
	message := fmt.Sprintf(`%d row affected`, data_row)

	return c.JSON(fiber.Map{
		"message": message,
		"data":    data_update,
	})
}
