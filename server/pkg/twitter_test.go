package tweetviz

import (
	"testing"

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
