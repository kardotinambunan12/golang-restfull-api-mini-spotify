package song

import (
	"final-project/app/services"
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func TestSongGetById(t *testing.T) {
	app := fiber.New()
	services.InitDatabaseForTest()

	app.Get("/song/:id", SongGetById)

	//Positive case
	req, _ := http.NewRequest("GET", "/song/1", nil)
	req.Header.Set("accept", "application/json")
	res, err := app.Test(req)

	utils.AssertEqual(t, nil, err, "send request")
	utils.AssertEqual(t, 200, res.StatusCode, "response success")

	// negative case
	// req, _ = http.NewRequest("GET", "/song/1", nil)
	// req.Header.Set("accept", "application/json")
	// res, err = app.Test(req)

	// utils.AssertEqual(t, nil, err, "send request")
	// utils.AssertEqual(t, 404, res.StatusCode, "bad request")

	req, _ = http.NewRequest("GET", "/song/1", nil)
	res, err = app.Test(req)

	utils.AssertEqual(t, nil, err, "send request")
	utils.AssertEqual(t, 400, res.StatusCode, "invalid accept header")

}
