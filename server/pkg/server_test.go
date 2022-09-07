package tweetviz

import (
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestSetupServer(t *testing.T) {
	server := SetupServer()
	assert.IsType(t, server, &fiber.App{}, "Could not set up fiber server.")
}
