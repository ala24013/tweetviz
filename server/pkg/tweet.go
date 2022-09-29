package tweetviz

import (
	"encoding/json"
	"errors"
	"fmt"

	twitter "github.com/g8rswimmer/go-twitter/v2"
)

type Tweet struct {
	Username string
	Tweet    string
	Geo      []float64
}

func (t Tweet) String() string {
	return fmt.Sprintf(
		"Username: %s\nTweet: %s\nGeo: [%f, %f]",
		t.Username, t.Tweet, t.Geo[1], t.Geo[0],
	)
}

func findCenter(box []float64) ([]float64, error) {
	if len(box) != 4 {
		return nil, errors.New("Invalid box, must contain 4 elements in slice")
	}
	longCenter := (box[0] + box[2]) / 2
	latCenter := (box[1] + box[3]) / 2
	return []float64{longCenter, latCenter}, nil
}

func processTweet(t *twitter.TweetMessage) (Tweet, error) {
	var aid string
	var username string
	var text string
	var centerGeo []float64
	for _, tweet := range t.Raw.Tweets {
		text = tweet.Text
		aid = tweet.AuthorID
	}
	for _, place := range t.Raw.Includes.Places {
		box := place.Geo.BBox
		center, err := findCenter(box)
		if err != nil {
			return Tweet{}, err
		}
		centerGeo = center
	}
	for _, user := range t.Raw.Includes.Users {
		if user.ID == aid {
			username = user.UserName
		}
	}

	// Invert from twitter's long/lat format to the map's lat/long format
	centerGeo[0], centerGeo[1] = centerGeo[1], centerGeo[0]

	tweet := Tweet{
		Username: username,
		Tweet:    text,
		Geo:      centerGeo,
	}

	return tweet, nil
}

func printRawTweet(t *twitter.TweetMessage) {
	enc, err := json.MarshalIndent(t, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(enc))
}

func printFormattedTweet(t Tweet) {
	fmt.Println(t.String())
}
