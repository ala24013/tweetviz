package tweetviz

import (
	"sync"

	"github.com/goccy/go-json"
	"golang.org/x/exp/maps"
)

// Tweetlist holds the list (map) of tweets and the associated lock.
// A map was chosen to represent the list rather than a slice for
// better efficiency adds/deletes at the cost of more expensive
// serializes
type Tweetlist struct {
	list map[string]Tweet
	lock sync.Mutex
}

// CreateTweetlist creates a new tweetlist
func CreateTweetlist() *Tweetlist {
	return &Tweetlist{
		list: map[string]Tweet{},
		lock: sync.Mutex{},
	}
}

// addTweet adds a tweet to the tweetlist
func (t *Tweetlist) addTweet(tweet Tweet) {
	// This key should be *mostly* unique and in the case it is not,
	// then it is boring
	key := tweet.Username + tweet.Tweet
	t.lock.Lock()
	t.list[key] = tweet
	t.lock.Unlock()
}

// delTweet removes a tweet from the tweetlist
func (t *Tweetlist) delTweet(tweet Tweet) {
	key := tweet.Username + tweet.Tweet
	t.lock.Lock()
	delete(t.list, key)
	t.lock.Unlock()
}

// serialize returns a slice of tweets for a tweetlist
func (t *Tweetlist) serialize() ([]byte, error) {
	vals := maps.Values(t.list)
	return json.Marshal(vals)
}
