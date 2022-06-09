package album

import (
	"bytes"
	"final-project/app/services"
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func TestAlbumPost(t *testing.T) {
	app := fiber.New()
	services.InitDatabaseForTest()
	app.Post("/album", AlbumPost)

	payload := bytes.NewReader([]byte(`
	{
		
			"album_name": "string",
			"artist_id": "ac8688d0-24b7-47bb-a181-eee664ec91e4",
			"gendre_id": "59d3c5bb-7e40-414a-a83c-2765daeaba54",
			"realese_date-date": "2022-06-08"
		  
	  }
	`))
	//positive case
	req, _ := http.NewRequest("POST", "/album", payload)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("accept", "application/json")
	res, err := app.Test(req)

	utils.AssertEqual(t, nil, err, "create request")
	utils.AssertEqual(t, 201, res.StatusCode, "Post success")

	// //negative case
	req, _ = http.NewRequest("POST", "/album", payload)
	res, err = app.Test(req)
	utils.AssertEqual(t, nil, err, "create request")
	utils.AssertEqual(t, 400, res.StatusCode, "invalid content-type")

	// negative case
	req, _ = http.NewRequest("POST", "/album", nil)
	req.Header.Set("Content-Type", "application/json")
	res, err = app.Test(req)
	utils.AssertEqual(t, nil, err, "create request")
	utils.AssertEqual(t, 400, res.StatusCode, "invalid payload")

	req, _ = http.NewRequest("POST", "/album", nil)
	req.Header.Set("Content-Type", "application/json")
	res, err = app.Test(req)
	utils.AssertEqual(t, nil, err, "create request")
	utils.AssertEqual(t, 400, res.StatusCode, "invalid payload")

	//negative case
	req, _ = http.NewRequest("POST", "/album", payload)
	res, err = app.Test(req)
	utils.AssertEqual(t, nil, err, "create request")
	utils.AssertEqual(t, 400, res.StatusCode, "invalid content-type")

	req, _ = http.NewRequest("POST", "/album", nil)
	req.Header.Set("accept", "application/json")
	res, err = app.Test(req)

	utils.AssertEqual(t, nil, err, "send request")
	utils.AssertEqual(t, 400, res.StatusCode, "post response body parser")

}
