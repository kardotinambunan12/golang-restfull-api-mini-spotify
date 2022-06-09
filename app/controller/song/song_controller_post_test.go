package song

import (
	"bytes"
	"final-project/app/services"
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func TestPostSong(t *testing.T) {
	app := fiber.New()
	services.InitDatabaseForTest()

	app.Post("/song", SongPost)

	payload := bytes.NewReader([]byte(`
	{
		"artist_id": "ac8688d0-24b7-47bb-a181-eee664ec91e4",
		"description": "Test",
		"gendre": "Test",
		"name": "Test",
		"year": 2010
	  }
	`))

	//positive case
	req, _ := http.NewRequest("POST", "/song", payload)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("accept", "application/json")
	res, err := app.Test(req)

	utils.AssertEqual(t, nil, err, "create request")
	utils.AssertEqual(t, 201, res.StatusCode, "Post success")

	//negative case
	req, _ = http.NewRequest("POST", "/song", payload)
	res, err = app.Test(req)
	utils.AssertEqual(t, nil, err, "create request")
	utils.AssertEqual(t, 400, res.StatusCode, "invalid content-type")

	req, _ = http.NewRequest("POST", "/song", nil)
	req.Header.Set("accept", "application/json")
	res, err = app.Test(req)

	utils.AssertEqual(t, nil, err, "send request")
	utils.AssertEqual(t, 400, res.StatusCode, "post response body parser")

}
