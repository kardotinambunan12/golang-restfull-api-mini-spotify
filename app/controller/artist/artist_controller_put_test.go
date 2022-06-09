package artist

import (
	"bytes"
	"final-project/app/model"
	"final-project/app/services"
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func TestArtistPut(t *testing.T) {

	// services.InitDatabaseForTest()
	db := services.DB
	app := fiber.New()

	app.Put("/artist/:id", ArtistPut)

	about := "Test"
	name := "Name Test"

	init := model.Artist{
		ArtistApi: model.ArtistApi{
			Name:  &name,
			About: &about,
		},
	}
	db.Create(&init)

	payload := bytes.NewReader([]byte(`
	{
		"about": "Test String",
		"name": "Test String"
	  }
	`))

	uri := "/artist/ac8688d0-24b7-47bb-a181-eee664ec91e4"

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

	req, _ = http.NewRequest("PUT", "/artist/90000", nil)
	req.Header.Set("Content-Type", "application/json")

	res, err = app.Test(req)
	utils.AssertEqual(t, nil, err, "send request")
	utils.AssertEqual(t, 400, res.StatusCode, "put response error accept header")

}
