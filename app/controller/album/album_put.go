package album

import (
	"final-project/app/model"
	"final-project/app/services"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// PutAlbum godoc
// @Summary Update Album by id
// @Description Update Album by id
// @Accept  application/json
// @Produce  application/json
// @Success 201 {object} model.AlbumApi "OK"
// @Success 401
// @Success 400
// @Param id path string true "Album ID"
// @Param data body model.AlbumApi true "Album data"
// @Router /album/{id} [put]
// @Tags Album
func AlbumPut(c *fiber.Ctx) error {
	db := services.DB
	var album model.AlbumApi

	accept := c.Get("accept")
	if !strings.EqualFold(accept, "application/json") {
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid accept header",
		})
	}

	id := c.Params("id")

	if err := c.BodyParser(&album); err != nil {
		fmt.Println("error :", err.Error())
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid request",
		})
	}

	updates := &model.Album{AlbumApi: album}
	result := db.Model(&model.Album{}).Where(`id=?`, id).Updates(updates)
	row_value := result.RowsAffected
	message := fmt.Sprintf(`%d row updated`, row_value)

	return c.JSON(fiber.Map{
		"message": message,
		"data":    updates,
	})
}
