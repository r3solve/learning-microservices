package main

import (
	"backend/helper"
	_ "fmt"
	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		help := `
			visit the /quote route to get quotes of the following categories
			usage = http://localhost:3000/quote?category=dad
			categories :)
			age
			alone
			amazing
			anger
			architecture
			art
			attitude
			beauty
			best
			birthday
			business
			car
			change
			communication
			computers
			cool
			courage
			dad
			dating
			death
			design
			dreams
			education
			environmental
			equality
			experience
			failure
			faith
			family
			famous
			fear
			fitness
			food
			forgiveness
			freedom
			friendship
			funny
			future
			god
			good
			government
			graduation
			great
			happiness
			health
			history
			home
			hope
			humor
			imagination
			inspirational
			intelligence
			jealousy
			knowledge
			leadership
			learning
			legal
			life
			love
			marriage
			medical
			men
			mom
			money
			morning
			movies
success
		`
		if err != nil {
			return nil
		}

		return c.SendString(help)
	})
	app.Get("/quote", func(c *fiber.Ctx) error {
		cat := c.Query("category")
		if cat == "" {
			return c.Status(fiber.StatusBadRequest).SendString("Category query parameter is required")
		}
		quote, _ := helper.GetQuote(cat)
		return c.JSON(map[string][]helper.Quote{"quote": quote})
	})

	app.Listen(":3000")
}
