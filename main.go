package main

import (
	"go-fiber-rest/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	// Use router movies
	routes.UseMoviesRoute(app)

	app.Listen(":4000")

}
