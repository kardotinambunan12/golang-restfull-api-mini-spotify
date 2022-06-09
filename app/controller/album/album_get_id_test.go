package album

import (
	"final-project/app/services"
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func TestAlbumGetById(t *testing.T) {
	app := fiber.New()
	services.InitDatabaseForTest()
	app.Get("/album/:id", AlbumGetById)

	//positif case
	req, _ := http.NewRequest("GET", "/album/1", nil)
	req.Header.Set("accept", "application/json")

	res, err := app.Test(req)
	utils.AssertEqual(t, nil, err, "send request")
	utils.AssertEqual(t, 200, res.StatusCode, "get response")

	//negative case
	req, _ = http.NewRequest("GET", "/album/1", nil)
	res, err = app.Test(req)
	utils.AssertEqual(t, nil, err, "send request")
	utils.AssertEqual(t, 400, res.StatusCode, "invalid accept && request")

}
