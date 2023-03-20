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

	router.Get("/:id", func(c *fiber.Ctx) error {

		id, _ := c.ParamsInt("id")

		var movieFound Movie

		for _, movie := range movies {
			if movie.Id == id {
				movieFound = *movie
			}
		}

		return c.JSON(fiber.Map{
			"movies": movieFound,
		})
	})

	router.Post("/", func(c *fiber.Ctx) error {

		type Request struct {
			Title string
			Id    int
		}
		var body Request
		c.BodyParser(&body)
		// fmt.Println(body)

		// CREAR NEW ITEM
		newMovie := Movie{
			Title: body.Title,
			Id:    len(movies) + 1,
		}
		// fmt.Println(newMovie)

		// AGREGAR EL NEW ITEM
		movies = append(movies, &newMovie)

		return c.JSON(fiber.Map{
			"movies": movies,
		})
	})

	router.Put("/:id", func(c *fiber.Ctx) error {

		// RECOGER EL PARAMETRO DEL ENDPOINT
		id, err := c.ParamsInt("id")

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "there's a error",
			})
		}

		// RECOGER EL BODY JSON
		type Request struct {
			Title string
			Id    int
		}
		var body Request
		c.BodyParser(&body)

		// REEMPLAZAR
		for _, movie := range movies {
			if movie.Id == id {
				movie.Title = body.Title
			}
		}

		return c.JSON(fiber.Map{
			"movies": movies,
		})
	})

	router.Delete("/:id", func(c *fiber.Ctx) error {

		// RECOGER EL PARAMETRO DEL ENDPOINT
		id, err := c.ParamsInt("id")

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "there's a error",
			})
		}

		// REEMPLAZAR
		for idx, movie := range movies {
			if movie.Id == id {
				movies = append(movies[:idx], movies[idx+1:]...)
			}
		}

		return c.JSON(fiber.Map{
			"movies": movies,
		})
	})

}
