package song

import (
	"final-project/app/services"
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func TestSongDeleted(t *testing.T) {

	app := fiber.New()
	services.InitDatabaseForTest()

	app.Delete("/song/:id", SongDelete)

	//positive case
	req, _ := http.NewRequest("DELETE", "/song/1", nil)
	req.Header.Set("accept", "application/json")
	res, err := app.Test(req)

	utils.AssertEqual(t, nil, err, "send request")
	utils.AssertEqual(t, 200, res.StatusCode, "deleted succes")

	// negative case
	// req, _ = http.NewRequest("DELETE", "/song/3000", nil)
	// req.Header.Set("accept", "application/json")
	// res, err = app.Test(req)

	// utils.AssertEqual(t, nil, err, "send request")
	// utils.AssertEqual(t, 404, res.StatusCode, "request not found")

	req, _ = http.NewRequest("DELETE", "/song/1", nil)

	res, err = app.Test(req)

	utils.AssertEqual(t, nil, err, "send request")
	utils.AssertEqual(t, 400, res.StatusCode, "invalid accept Header")

}
