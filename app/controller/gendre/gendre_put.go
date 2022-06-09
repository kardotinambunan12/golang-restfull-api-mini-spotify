package gendre

import (
	"final-project/app/model"
	"final-project/app/services"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// PutGendre Go doc
// @Summary Updated Gendre by Id
// @Description udpated Gendre by Id
// @Accept application/json
// @Produce application/json
// @Success 201 {object} model.GendreApi "OK"
// @Success 401
// @Success 404
// @Param id path string true "Gendre ID"
// @Param data body model.GendreApi true "Gendre data"
// @Router /gendre/{id} [put]
// @Tags Gendre
func GendrePut(c *fiber.Ctx) error {

	db := services.DB

	accept := c.Get("accept")

	if !strings.EqualFold(accept, "application/json") {
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid headers",
		})
	}
	var gendre model.GendreApi
	id := c.Params("id")

	if err := c.BodyParser(&gendre); err != nil {
		fmt.Println("error :", err.Error())
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid request",
		})
	}

	data_update := &model.Gendre{GendreApi: gendre}
	result := db.Model(&model.Gendre{}).Where(`id=?`, id).Updates(data_update)
	data_row := result.RowsAffected
	msg := fmt.Sprintf(`%d row affected`, data_row)

	return c.JSON(fiber.Map{
		"message": msg,
		"status":  201,
		"data":    data_update,
	})

}
