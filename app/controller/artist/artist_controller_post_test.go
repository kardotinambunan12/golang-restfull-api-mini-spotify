package artist

import (
	"bytes"
	"final-project/app/services"
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func TestArtistPost(t *testing.T) {

	app := fiber.New()
	services.InitDatabaseForTest()
	app.Post("/artist", ArtistPost)

	payload := bytes.NewReader([]byte(`
	{
		"about": "Test",
		"name": "Post"
	  }
	`))
	//positive case
	req, _ := http.NewRequest("POST", "/artist", payload)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("accept", "application/json")
	res, err := app.Test(req)

	utils.AssertEqual(t, nil, err, "create request")
	utils.AssertEqual(t, 201, res.StatusCode, "Post success")

	//negative case
	req, _ = http.NewRequest("POST", "/artist", payload)
	res, err = app.Test(req)
	utils.AssertEqual(t, nil, err, "create request")
	utils.AssertEqual(t, 400, res.StatusCode, "invalid content-type")

	//negative case
	req, _ = http.NewRequest("POST", "/artist", nil)
	req.Header.Set("Content-Type", "application/json")
	res, err = app.Test(req)
	utils.AssertEqual(t, nil, err, "create request")
	utils.AssertEqual(t, 400, res.StatusCode, "invalid payload")

}
