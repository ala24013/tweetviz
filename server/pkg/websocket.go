package tweetviz

import (
	"log"

	"github.com/gofiber/websocket/v2"
)

// client represents a websocket client
type client struct{}

var clients = make(map[*websocket.Conn]client)
var register = make(chan *websocket.Conn)
var fromClient = make(chan []byte)
var toClient = make(chan []byte)
var unregister = make(chan *websocket.Conn)

// runWebsockets runs the communications between the server and the client
// using the websocket.
func runWebsockets() {
	for {
		select {
		case c := <-register:
			clients[c] = client{}
			log.Println("connection registered")

		case message := <-fromClient:
			m := string(message)
			log.Println("message received:", m)

		case tweetlist := <-toClient:
			for c := range clients {
				if err := c.WriteMessage(websocket.TextMessage, tweetlist); err != nil {
					log.Println("write error:", err)

					unregister <- c
					c.WriteMessage(websocket.CloseMessage, []byte{})
					c.Close()
				}
			}

		case c := <-unregister:
			delete(clients, c)
			log.Println("connection unregistered")
		}
	}
}
