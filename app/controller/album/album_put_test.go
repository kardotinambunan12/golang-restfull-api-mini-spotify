package album

import (
	"bytes"
	// "final-project/app/model"
	// "final-project/app/services"
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func TestAlbumPut(t *testing.T) {

	// services.InitDatabaseForTest()
	// db := services.DB
	app := fiber.New()

	app.Put("/album/:id", AlbumPut)

	payload := bytes.NewReader([]byte(`
	{
		"album_name": "string",
		"artist_id": "ac8688d0-24b7-47bb-a181-eee664ec91e4",
		"gendre_id": "59d3c5bb-7e40-414a-a83c-2765daeaba54",
		"realese_date-date": "2022-06-08"
	  }
	`))

	uri := "/album/5b142996-6148-44bc-88ee-d84d9fb2af18"

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

	req, _ = http.NewRequest("PUT", "/album/000000", nil)
	req.Header.Set("Content-Type", "application/json")

	res, err = app.Test(req)
	utils.AssertEqual(t, nil, err, "send request")
	utils.AssertEqual(t, 400, res.StatusCode, "put response error accept header")

}
