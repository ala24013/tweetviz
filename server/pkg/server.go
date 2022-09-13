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

/*
type client struct{}

var clients = make(map[*websocket.Conn]client)
var  = make(chan *websocket.Conn)

func runBroker(tweets chan *) {
	for c := range tweets {
		for connection := range clients {
			if err := connection.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
				log.Println("write error:", err)

				connection.WriteMessage(websocket.CloseMessage, []byte{})
				connection.Close()
				delete(clients, connection)
			}
		}
	}
}*/

func SetupServer() *fiber.App {
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
			msg := time.Now()
			err := c.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("%d", msg)))
			if err != nil {
				log.Println("write:", err)
				break
			}
		}
	}))

	return app
}

func RunServer() {
	server := SetupServer()
	log.Fatal(server.Listen(":3000"))
}
