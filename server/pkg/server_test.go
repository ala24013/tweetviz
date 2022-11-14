package tweetviz

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestSetupServer(t *testing.T) {
	tl := CreateTweetlist()
	server := SetupServer(tl)
	assert.IsType(t, server, &fiber.App{}, "Could not set up fiber server.")
}

func TestHelloRoute(t *testing.T) {
	tl := CreateTweetlist()
	app := SetupServer(tl)
	req := httptest.NewRequest("GET", "/hello", nil)
	resp, _ := app.Test(req, 1)
	assert.Equalf(t, 200, resp.StatusCode, "failed to get proper response for hello route")
}

func TestWsRoute(t *testing.T) {
	tl := CreateTweetlist()
	app := SetupServer(tl)
	req := httptest.NewRequest("GET", "/ws", nil)
	resp, _ := app.Test(req, 1)
	assert.Equalf(t, 426, resp.StatusCode, "failed to get proper response for ws route")
}
