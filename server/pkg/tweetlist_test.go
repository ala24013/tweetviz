package tweetviz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTweetlist(t *testing.T) {
	tl := CreateTweetlist()
	assert.IsType(t, tl, &Tweetlist{}, "could not get create tweetlist")
}

func TestAddTweet(t *testing.T) {
	tw := Tweet{Username: "Bob", Tweet: "Hello, World", Geo: []float64{1.1, 2.2}}
	tl := CreateTweetlist()
	tl.addTweet(tw)
	assert.Equal(t, tl.list["BobHello, World"], tw)
}

func TestDelTweet(t *testing.T) {
	tw := Tweet{Username: "Bob", Tweet: "Hello, World", Geo: []float64{1.1, 2.2}}
	tl := CreateTweetlist()
	tl.addTweet(tw)
	tl.delTweet(tw)
	assert.Equal(t, len(tl.list), 0)
}

func TestSerialize(t *testing.T) {
	expected := []byte("[{\"Username\":\"Bob\",\"Tweet\":\"Hello, World\",\"Geo\":[1.1,2.2]}]")
	tw := Tweet{Username: "Bob", Tweet: "Hello, World", Geo: []float64{1.1, 2.2}}
	tl := CreateTweetlist()
	tl.addTweet(tw)
	s, err := tl.serialize()
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, s, expected)
}
