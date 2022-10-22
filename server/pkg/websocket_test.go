package tweetviz

import (
	"testing"

	"github.com/gofiber/websocket/v2"
)

func TestRunWebsockets(t *testing.T) {
	c := &websocket.Conn{}
	tl := CreateTweetlist()
	go runWebsockets(tl)
	register <- c
	unregister <- c
	// Skipping other channels because the effects would break our fake websocket
	shutdownWs <- 1
}
