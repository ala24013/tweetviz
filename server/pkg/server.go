package tweetviz

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/websocket/v2"
)

// SetupServer creates the fiber application
func SetupServer(t *Tweetlist) *fiber.App {
	app := fiber.New()

	app.Use(cors.New())
	app.Use(recover.New())

	go runWebsockets()
	go runTweetlist(t)

	app.Static("/", "../client/build")
	app.Static("/logo.png", "../client/build/logo.png")
	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/ws", websocket.New(func(c *websocket.Conn) {
		defer func() {
			unregister <- c
			c.Close()
		}()

		register <- c

		for {
			mt, m, err := c.ReadMessage()
			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					log.Println("read error:", err)
				}
				return
			}

			if mt == websocket.TextMessage {
				fromClient <- m
			} else {
				log.Println("websocket message received of type", mt)
			}
		}
	}))

	return app
}

// RunServer runs the webserver
func RunServer(t *Tweetlist) {
	server := SetupServer(t)
	log.Fatal(server.Listen("localhost:3000"))
}
