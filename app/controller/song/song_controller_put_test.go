package song

import (
	"bytes"
	"final-project/app/services"
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func TestPutSong(t *testing.T) {

	app := fiber.New()
	services.InitDatabaseForTest()

	app.Put("/artist/:id", SongPut)

	payload := bytes.NewReader([]byte(`
	{
		"artist_id": "ac8688d0-24b7-47bb-a181-eee664ec91e4",
		"description": "Test",
		"gendre": "Test",
		"name": "Test",
		"year": 2010
	  }
	`))

	uri := "/artist/24c71059-a983-4784-b8d0-f6c724f62272"

	//positive case
	req, _ := http.NewRequest("PUT", uri, payload)
	req.Header.Set("accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	res, err := app.Test(req)
	utils.AssertEqual(t, nil, err, "send request")
	utils.AssertEqual(t, 200, res.StatusCode, "put response success")

	// negative case
	req, _ = http.NewRequest("PUT", uri, nil)
	req.Header.Set("accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	res, err = app.Test(req)
	utils.AssertEqual(t, nil, err, "send request")
	utils.AssertEqual(t, 400, res.StatusCode, "put response body parser")

	req, _ = http.NewRequest("PUT", uri, nil)
	req.Header.Set("Content-Type", "application/json")

	res, err = app.Test(req)
	utils.AssertEqual(t, nil, err, "send request")
	utils.AssertEqual(t, 400, res.StatusCode, "invalid accept header")

}
