package tweetviz

import (
	"testing"

	twitter "github.com/g8rswimmer/go-twitter/v2"
	"github.com/stretchr/testify/assert"
)

func generateTestTweet(passing bool) *twitter.TweetMessage {
	bbox := []float64{1, 2, 3}
	if passing == true {
		bbox = []float64{1, 2, 3, 4}
	}
	tw := &twitter.TweetMessage{
		&twitter.TweetRaw{
			Tweets: []*twitter.TweetObj{
				&twitter.TweetObj{
					Text:     "hello",
					AuthorID: "12345",
				},
			},
			Includes: &twitter.TweetRawIncludes{
				Places: []*twitter.PlaceObj{
					&twitter.PlaceObj{
						Geo: &twitter.PlaceGeoObj{
							BBox: bbox,
						},
					},
				},
				Users: []*twitter.UserObj{
					&twitter.UserObj{
						UserName: "testuser",
					},
				},
			},
		},
	}
	return tw
}

func TestTweetString(t *testing.T) {
	tweet := Tweet{
		Username: "hi",
		Tweet:    "hello",
		Geo:      []float64{0, 1},
	}
	r := tweet.String()
	if r != "Username: hi\nTweet: hello\nGeo: [1.000000, 0.000000]" {
		t.Errorf("Did not properly stringify tweet.")
	}
}

func TestFindCenter(t *testing.T) {
	box := []float64{0, 1, 2, 3}
	r, err := findCenter(box)
	if err != nil {
		t.Errorf("Errored when trying to find center")
	}
	if r[0] != 1 {
		t.Errorf("Did not properly average longitudes")
	}
	if r[1] != 2 {
		t.Errorf("Did not properly average longitudes")
	}
}
func TestFailFindCenter(t *testing.T) {
	box := []float64{1, 2, 3}
	_, err := findCenter(box)
	if err == nil {
		t.Errorf("Should fail when box does not have 4 elements")
	}
}

func TestProcessTweet(t *testing.T) {
	tw := generateTestTweet(true)
	tweet, err := processTweet(tw)
	if err != nil {
		t.Errorf("Did not process tweet correctly: %v", err)
	}
	assert.IsTypef(t, Tweet{}, tweet, "did not process tweet correctly: bad type")
}

func TestFailProcessTweet(t *testing.T) {
	tw := generateTestTweet(false)
	_, err := processTweet(tw)
	if err == nil {
		t.Errorf("Did not process tweet correctly: %v", err)
	}
}
