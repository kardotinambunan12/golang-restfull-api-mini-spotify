package playlist

import (
	"bytes"
	"final-project/app/services"
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func TestPlaylistPost(t *testing.T) {
	app := fiber.New()
	services.InitDatabaseForTest()
	app.Post("/playlist", PlaylistPost)

	payload := bytes.NewReader([]byte(`
	{
		"creation_date": "2010-06-08",
		"playlist_desc": "my playlist descriptions",
		"playlist_name": "Sintrong",
		"playlist_pic": "picture AMD 64",
		"playlist_type_id": "b7e45ac4-028e-4f60-892b-001794fcca2e"
	  }
	`))
	//positive case
	req, _ := http.NewRequest("POST", "/playlist", payload)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("accept", "application/json")
	res, err := app.Test(req)

	utils.AssertEqual(t, nil, err, "create request")
	utils.AssertEqual(t, 201, res.StatusCode, "Post success")

	// //negative case
	req, _ = http.NewRequest("POST", "/playlist", payload)
	res, err = app.Test(req)
	utils.AssertEqual(t, nil, err, "create request")
	utils.AssertEqual(t, 400, res.StatusCode, "invalid content-type")

	// negative case
	req, _ = http.NewRequest("POST", "/playlist", nil)
	req.Header.Set("Content-Type", "application/json")
	res, err = app.Test(req)
	utils.AssertEqual(t, nil, err, "create request")
	utils.AssertEqual(t, 400, res.StatusCode, "invalid payload")

	req, _ = http.NewRequest("POST", "/playlist", nil)
	req.Header.Set("Content-Type", "application/json")
	res, err = app.Test(req)
	utils.AssertEqual(t, nil, err, "create request")
	utils.AssertEqual(t, 400, res.StatusCode, "invalid payload")

	//negative case
	req, _ = http.NewRequest("POST", "/playlist", payload)
	res, err = app.Test(req)
	utils.AssertEqual(t, nil, err, "create request")
	utils.AssertEqual(t, 400, res.StatusCode, "invalid content-type")

	req, _ = http.NewRequest("POST", "/playlist", nil)
	req.Header.Set("accept", "application/json")
	res, err = app.Test(req)

	utils.AssertEqual(t, nil, err, "send request")
	utils.AssertEqual(t, 400, res.StatusCode, "post response body parser")
}
