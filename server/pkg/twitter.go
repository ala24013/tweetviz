package tweetviz

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/coreos/pkg/flagutil"
	twitter "github.com/g8rswimmer/go-twitter/v2"
)

var (
	TWITTER_BEARER_TOKEN string
)

// Stream conducts the high level logic to manage the rules and stream tweets
func Stream(query string, t *Tweetlist) {
	deleteAllRules()
	addRule(query)
	streamTweets(t)
}

// wrappedStream wraps the TweetSearchStream function, attempting to
// connect a stream. Since twitter is lazy in checking whether a stream
// is alive or not, it will be slow to update that an old client has
// disconnected, and will frequently be overly aggressive with its
// TooManyConnections errors.
func wrappedStream(client *twitter.Client, opts twitter.TweetSearchStreamOpts) (*twitter.TweetStream, error) {
	for i := 1; i < 20; i++ {
		log.Printf("Attempting to connect stream.")
		s, err := client.TweetSearchStream(context.Background(), opts)
		if err != nil {
			e, ok := err.(*twitter.ErrorResponse)
			if ok {
				if e.StatusCode == 429 {
					// Client is connected error, let's wait 1 second and try again
					time.Sleep(1 * time.Second)
					continue
				}
			}
			log.Panicf("%v", err)
		}
		log.Printf("Stream connected.")
		doneLoading <- 0
		return s, nil
	}
	return nil, errors.New("failed connecting 20 times")
}

// streamTweets streams in the tweets from Twitter
func streamTweets(t *Tweetlist) {
	client, err := getClient()
	if err != nil {
		panic(err)
	}

	opts := twitter.TweetSearchStreamOpts{
		Expansions: []twitter.Expansion{
			twitter.ExpansionAuthorID,
			twitter.ExpansionGeoPlaceID,
		},
		PlaceFields: []twitter.PlaceField{
			twitter.PlaceFieldGeo,
		},
		UserFields: []twitter.UserField{
			twitter.UserFieldName,
			twitter.UserFieldID,
		},
	}

	s, err := wrappedStream(client, opts)
	if err != nil {
		log.Panicf("%v", err)
	}

	func() {
		defer s.Close()
		for {
			select {
			case <-streamShutdown:
				log.Println("closing stream")
				return
			case tm := <-s.Tweets():
				tw, err := processTweet(tm)
				if err != nil {
					log.Printf("error decoding tweet %v", err)
				}
				t.addTweet(tw)
				go func() {
					time.Sleep(8 * time.Second)
					t.delTweet(tw)
				}()
			case sm := <-s.SystemMessages():
				smb, err := json.Marshal(sm)
				if err != nil {
					log.Printf("error decoding system message %v", err)
				}
				log.Printf("system: %s\n\n", string(smb))
			case strErr := <-s.Err():
				log.Printf("error: %v\n\n", strErr)
			default:
			}
			if s.Connection() == false {
				log.Println("connection lost")
				return
			}
		}
	}()
}

// addRule adds a new rule to the twitter stream
func addRule(query string) {
	client, err := getClient()
	if err != nil {
		panic(err)
	}

	streamRule := twitter.TweetSearchStreamRule{
		Value: fmt.Sprintf("%s has:geo", query),
		Tag:   fmt.Sprintf("%s rule", query),
	}

	_, err = client.TweetSearchStreamAddRule(context.Background(), []twitter.TweetSearchStreamRule{streamRule}, false)
	if err != nil {
		log.Panicf("tweet search stream add rule callout error: %v", err)
	}
}

// getRules returns the slice of twitter rules that are currently in place on the stream
func getRules() []*twitter.TweetSearchStreamRuleEntity {
	client, err := getClient()
	if err != nil {
		panic(err)
	}

	searchStreamRules, err := client.TweetSearchStreamRules(context.Background(), []twitter.TweetSearchStreamRuleID{})
	if err != nil {
		log.Panicf("tweet search stream rule callout error: %v", err)
	}

	return searchStreamRules.Rules
}

// deleteAllRules removes all the existing rules on the twitter stream
func deleteAllRules() {
	rules := getRules()
	ruleIds := make([]string, len(rules))
	for i, r := range rules {
		ruleIds[i] = string(r.ID)
	}
	deleteRules(ruleIds)
}

// deleteRules removes rules by their id's from the twitter stream
func deleteRules(ids []string) {
	if len(ids) > 0 {
		client, err := getClient()
		if err != nil {
			panic(err)
		}

		ruleIDs := []twitter.TweetSearchStreamRuleID{}
		for _, id := range ids {
			ruleIDs = append(ruleIDs, twitter.TweetSearchStreamRuleID(id))
		}

		_, err = client.TweetSearchStreamDeleteRuleByID(context.Background(), ruleIDs, false)
		if err != nil {
			log.Panicf("tweet search stream delete rule callout error: %v", err)
		}
	}
}

// authorize is the scructure containing the twitter bearer token
type authorize struct {
	Token string
}

// Add adds an authorization header to an HTTP request
func (a authorize) Add(req *http.Request) {
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", a.Token))
}

// getClient gets a Twitter client
func getClient() (*twitter.Client, error) {
	if TWITTER_BEARER_TOKEN == "" {
		flag.StringVar(&TWITTER_BEARER_TOKEN, "bearer-token", "", "Twitter Bearer Token")
	}
	flag.Parse()
	err := flagutil.SetFlagsFromEnv(flag.CommandLine, "TWITTER")
	if err != nil {
		panic("TWITTER_BEARER_TOKEN Not found!")
	}

	client := &twitter.Client{
		Authorizer: authorize{
			Token: TWITTER_BEARER_TOKEN,
		},
		Client: http.DefaultClient,
		Host:   "https://api.twitter.com",
	}

	return client, nil
}
