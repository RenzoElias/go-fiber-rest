package routes

import (
	"github.com/gofiber/fiber/v2"
)

type Movie struct {
	Title string `json: "title"`
	Id    int    `json: "id"`
}

func UseMoviesRoute(router fiber.Router) {

	movies := []*Movie{
		{
			Title: "kimetsu no yaiba",
			Id:    1,
		},
		{
			Title: "movies 2",
			Id:    2,
		},
	}

	router.Get("/", func(c *fiber.Ctx) error {
		// return c.SendString("Hello from fiber")
		return c.JSON(fiber.Map{
			"movies": movies,
		})
	})
}
