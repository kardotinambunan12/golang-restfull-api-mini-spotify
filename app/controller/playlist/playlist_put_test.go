package playlist

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func TestPlaylistPut(t *testing.T) {
	app := fiber.New()

	app.Put("/playlist/:id", PlaylistPut)

	payload := bytes.NewReader([]byte(`
	{
		"creation_date": "2012-06-08",
		"playlist_desc": "Playlist Description",
		"playlist_name": "Sintrong",
		"playlist_pic": "Image Picture AMD 64",
		"playlist_type_id": "b7e45ac4-028e-4f60-892b-001794fcca2e"
	  }
	`))

	uri := "/playlist/5b142996-6148-44bc-88ee-d84d9fb2af18"

	//positive case
	req, _ := http.NewRequest("PUT", uri, payload)
	req.Header.Set("accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	res, err := app.Test(req)
	utils.AssertEqual(t, nil, err, "send request")
	utils.AssertEqual(t, 200, res.StatusCode, "put response success")

	//negative case
	req, _ = http.NewRequest("PUT", uri, nil)
	req.Header.Set("Content-Type", "application/json")

	res, err = app.Test(req)
	utils.AssertEqual(t, nil, err, "send request")
	utils.AssertEqual(t, 400, res.StatusCode, "put response error payload header")

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
	utils.AssertEqual(t, 400, res.StatusCode, "put response error accept header")

	req, _ = http.NewRequest("PUT", "/playlist/000000", nil)
	req.Header.Set("Content-Type", "application/json")

	res, err = app.Test(req)
	utils.AssertEqual(t, nil, err, "send request")
	utils.AssertEqual(t, 400, res.StatusCode, "put response error accept header")
}
