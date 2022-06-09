package gendre

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func TestGendrePut(t *testing.T) {
	// services.InitDatabaseForTest()
	// db := services.DB
	app := fiber.New()

	app.Put("/gendre/:id", GendrePut)

	payload := bytes.NewReader([]byte(`
	{
		"genreDescr": "string",
		"genreName": "string"
	  }
	`))

	uri := "/gendre/5b142996-6148-44bc-88ee-d84d9fb2af18"

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

}
