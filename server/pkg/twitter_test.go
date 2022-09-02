package tweetviz

import (
	"testing"
	"time"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/stretchr/testify/assert"
)

func TestGetClient(t *testing.T) {
	client, err := getClient()
	if err != nil {
		t.Errorf("Error getting twitter client: %v", err)
	}
	assert.IsType(t, client, &twitter.Client{}, "Could not get twitter client.")
}

func TestStream(t *testing.T) {
	shutdown := make(chan int)
	query := "hello"
	go Stream(query, shutdown)
	time.Sleep(1 * time.Second)
	shutdown <- 1
}
