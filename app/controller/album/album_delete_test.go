package album

import (
	"final-project/app/services"
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func TestAlbumDelete(t *testing.T) {

	app := fiber.New()
	services.InitDatabaseForTest()

	app.Delete("/album/:id", AlbumDelete)

	//positif case
	req, _ := http.NewRequest("DELETE", "/album/1", nil)
	req.Header.Set("accept", "application/json")
	res, err := app.Test(req)
	utils.AssertEqual(t, nil, err, "send request")
	utils.AssertEqual(t, 200, res.StatusCode, "get response 200")

	//negatife case
	req, _ = http.NewRequest("DELETE", "/album/500", nil)
	req.Header.Set("accept", "application/json")
	res, err = app.Test(req)
	utils.AssertEqual(t, nil, err, "send request")
	utils.AssertEqual(t, 200, res.StatusCode, "request not found")

	//negatife case
	req, _ = http.NewRequest("DELETE", "/album/1", nil)
	res, err = app.Test(req)
	utils.AssertEqual(t, nil, err, "send request")
	utils.AssertEqual(t, 400, res.StatusCode, "invalid accept")

}
