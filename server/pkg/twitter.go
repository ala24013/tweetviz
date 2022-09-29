package tweetviz

import (
	"context"
	"encoding/json"
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

func Stream(query string, t *Tweetlist, shutdown <-chan int) {
	deleteAllRules()
	addRule(query)
	streamTweets(t, shutdown)
}

func streamTweets(t *Tweetlist, shutdown <-chan int) {
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

	fmt.Println(opts)
	s, err := client.TweetSearchStream(context.Background(), opts)
	if err != nil {
		log.Panicf("tweet sample callout error: %v", err)
	}

	func() {
		defer s.Close()
		for {
			select {
			case <-shutdown:
				fmt.Println("closing")
				return
			case tm := <-s.Tweets():
				tw, err := processTweet(tm)
				if err != nil {
					fmt.Printf("error decoding tweet %v", err)
				}
				t.addTweet(tw)
				go func() {
					time.Sleep(5 * time.Second)
					t.delTweet(tw)
				}()
				//printFormattedTweet(tw)
			case sm := <-s.SystemMessages():
				smb, err := json.Marshal(sm)
				if err != nil {
					fmt.Printf("error decoding system message %v", err)
				}
				fmt.Printf("system: %s\n\n", string(smb))
			case strErr := <-s.Err():
				fmt.Printf("error: %v\n\n", strErr)
			default:
			}
			if s.Connection() == false {
				fmt.Println("connection lost")
				return
			}
		}
	}()
}

func addRule(query string) {
	client, err := getClient()
	if err != nil {
		panic(err)
	}

	streamRule := twitter.TweetSearchStreamRule{
		Value: fmt.Sprintf("%s has:geo", query),
		Tag:   fmt.Sprintf("%s rule", query),
	}

	searchStreamRules, err := client.TweetSearchStreamAddRule(context.Background(), []twitter.TweetSearchStreamRule{streamRule}, false)
	if err != nil {
		log.Panicf("tweet search stream add rule callout error: %v", err)
	}

	enc, err := json.MarshalIndent(searchStreamRules, "", "    ")
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(string(enc))
}

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

func deleteAllRules() {
	rules := getRules()
	ruleIds := make([]string, len(rules))
	for i, r := range rules {
		ruleIds[i] = string(r.ID)
	}
	deleteRules(ruleIds)
}

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
		searchStreamRules, err := client.TweetSearchStreamDeleteRuleByID(context.Background(), ruleIDs, false)
		if err != nil {
			log.Panicf("tweet search stream delete rule callout error: %v", err)
		}

		enc, err := json.MarshalIndent(searchStreamRules, "", "    ")
		if err != nil {
			log.Panic(err)
		}
		fmt.Println(string(enc))
	}
}

type authorize struct {
	Token string
}

func (a authorize) Add(req *http.Request) {
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", a.Token))
}

func getClient() (*twitter.Client, error) {
	//token := flag.String("token", os.Getenv("TWITTER_BEARER_TOKEN"), "twitter API token")
	if TWITTER_BEARER_TOKEN == "" {
		flag.StringVar(&TWITTER_BEARER_TOKEN, "bearer-token", "", "Twitter Bearer Token")
	}
	flag.Parse()
	flagutil.SetFlagsFromEnv(flag.CommandLine, "TWITTER")

	client := &twitter.Client{
		Authorizer: authorize{
			Token: TWITTER_BEARER_TOKEN,
		},
		Client: http.DefaultClient,
		Host:   "https://api.twitter.com",
	}

	return client, nil
}
