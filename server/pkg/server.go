package tweetviz

import (
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/websocket/v2"
)

func SetupServer(t *Tweetlist) *fiber.App {
	app := fiber.New()

	app.Use(cors.New())
	app.Use(recover.New())

	app.Static("/", "../client/build")
	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	app.Get("/ws", websocket.New(func(c *websocket.Conn) {
		fmt.Println("HELLO")
		for {
			time.Sleep(1 * time.Second)
			s, err := t.serialize()
			if err != nil {
				log.Println("Failed to serialize!")
				break
			}
			err = c.WriteMessage(websocket.TextMessage, s)
			if err != nil {
				log.Println("write:", err)
				break
			}
		}
	}))

	return app
}

func RunServer(t *Tweetlist) {
	server := SetupServer(t)
	log.Fatal(server.Listen(":3000"))
}
