package tweetviz

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func SetupServer() *fiber.App {
	app := fiber.New()

	app.Use(cors.New())
	app.Use(recover.New())

	app.Static("/", "../client/build")
	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	return app
}

func RunServer() {
	server := SetupServer()
	log.Fatal(server.Listen(":3000"))
}
