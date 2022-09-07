package tweetviz

import (
	"testing"

	twitter "github.com/g8rswimmer/go-twitter/v2"
	"github.com/stretchr/testify/assert"
)

func TestGetClient(t *testing.T) {
	client, err := getClient()
	if err != nil {
		t.Errorf("Error getting twitter client: %v", err)
	}
	assert.IsType(t, client, &twitter.Client{}, "Could not get twitter client.")
}

/*func TestStream(t *testing.T) {
	shutdown := make(chan int)
	query := "hello"
	go Stream(query, shutdown)
	time.Sleep(1 * time.Second)
	shutdown <- 1
}*/
