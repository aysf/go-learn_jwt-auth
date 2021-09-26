package main

import (
	"github.com/aysf/go_learn-jw_auth/database"
	"github.com/aysf/go_learn-jw_auth/routes"
	fiber "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	database.Connect()

	app := fiber.New()
	routes.Setup(app)

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	app.Listen(":8000")
}
