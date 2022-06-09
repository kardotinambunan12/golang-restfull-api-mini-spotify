package main

import (
	"final-project/app/config"
	"final-project/app/lib"
	"final-project/app/routes"
	"final-project/app/services"

	// "api/app/config"
	// "api/app/lib"
	// "api/app/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func init() {
	lib.LoadEnvironment(config.Environment)
}

// @title My Project
// @version 1.0.0
// @description API Documentation
// @termsOfService https://dospecs.monstercode.net/en/guide/tnc.html
// @contact.name Developer
// @contact.email developer.team.tog@gmail.com
// @host localhost:8000
// @schemes http
// @BasePath /api/v1

func main() {
	services.InitDatabase()

	app := fiber.New(fiber.Config{
		Prefork: viper.GetString("PREFORK") == "true",
	})

	routes.Handle(app)
	log.Fatal(app.Listen(":" + viper.GetString("PORT")))
}
