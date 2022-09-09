package tweetviz

import (
	"testing"
)

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
