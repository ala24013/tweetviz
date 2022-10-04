package tweetviz

import (
	"sync"

	"github.com/goccy/go-json"
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

// Values returns the values of the map m.
// The values will be in an indeterminate order.
// This function was datekn directly from golang.org/x/exp/maps .
// The only reason it is here is that the maps package is not
// currently supported in Golang 1.19
func Values[M ~map[K]V, K comparable, V any](m M) []V {
	r := make([]V, 0, len(m))
	for _, v := range m {
		r = append(r, v)
	}
	return r
}

// serialize returns a slice of tweets for a tweetlist
func (t *Tweetlist) serialize() ([]byte, error) {
	vals := Values(t.list)
	return json.Marshal(vals)
}
