package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type Movie struct {
	Title string
	Id    int
}

func main() {
	fmt.Println("Hello world")

	app := fiber.New()

	movies := []*Movie{
		{
			Title: "kimetsu no yaiba",
			Id:    1,
		},
	}

	app.Get("/", func(c *fiber.Ctx) error {
		// return c.SendString("Hello from fiber")
		return c.JSON(fiber.Map{
			"movies": movies,
		})
	})

	app.Listen(":4000")
}
