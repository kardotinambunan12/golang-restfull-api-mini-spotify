package gendre

import (
	"bytes"
	"final-project/app/services"
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func TestGendrePost(t *testing.T) {
	app := fiber.New()
	services.InitDatabaseForTest()
	app.Post("/gendre", GendrePost)

	payload := bytes.NewReader([]byte(`
	{
			
		"genreDescr": "string",
		"genreName": "string"
	  
}
	`))

	//positive case
	req, _ := http.NewRequest("POST", "/gendre", payload)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("accept", "application/json")
	res, err := app.Test(req)

	utils.AssertEqual(t, nil, err, "create request")
	utils.AssertEqual(t, 201, res.StatusCode, "Post success")

	// //negative case
	req, _ = http.NewRequest("POST", "/gendre", payload)
	res, err = app.Test(req)
	utils.AssertEqual(t, nil, err, "create request")
	utils.AssertEqual(t, 400, res.StatusCode, "invalid content-type")

	// negative case
	req, _ = http.NewRequest("POST", "/gendre", nil)
	req.Header.Set("Content-Type", "application/json")
	res, err = app.Test(req)
	utils.AssertEqual(t, nil, err, "create request")
	utils.AssertEqual(t, 400, res.StatusCode, "invalid payload")

	//negative case
	req, _ = http.NewRequest("POST", "/gendre", payload)
	res, err = app.Test(req)
	utils.AssertEqual(t, nil, err, "create request")
	utils.AssertEqual(t, 400, res.StatusCode, "invalid content-type")

	req, _ = http.NewRequest("POST", "/gendre", nil)
	req.Header.Set("accept", "application/json")
	res, err = app.Test(req)

	utils.AssertEqual(t, nil, err, "send request")
	utils.AssertEqual(t, 400, res.StatusCode, "post response body parser")

}
