package album

import (
	"final-project/app/services"
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func TestAlbumGetAll(t *testing.T) {
	app := fiber.New()

	services.InitDatabaseForTest()
	app.Get("/album", AlbumGetAll)

	// positive case
	req, _ := http.NewRequest("GET", "/album", nil)
	req.Header.Set("accept", "application/json")
	res, err := app.Test(req)

	utils.AssertEqual(t, nil, err, "send request")
	utils.AssertEqual(t, 200, res.StatusCode, "response success")

	//negative case
	req, _ = http.NewRequest("GET", "/album", nil)
	res, err = app.Test(req)

	utils.AssertEqual(t, nil, err, "send request")
	utils.AssertEqual(t, 400, res.StatusCode, "response success")
}
