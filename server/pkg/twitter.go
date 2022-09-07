package tweetviz

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/coreos/pkg/flagutil"
	twitter "github.com/g8rswimmer/go-twitter/v2"
)

var (
	TWITTER_BEARER_TOKEN string
)

func Stream(query string, shutdown <-chan int) {
	addRule(query)
	getRules()
	time.Sleep(1 * time.Second)
	<-shutdown
}

func addRule(query string) {
	client, err := getClient()
	if err != nil {
		panic(err)
	}

	streamRule := twitter.TweetSearchStreamRule{
		Value: query,
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

func getRules() {
	client, err := getClient()
	if err != nil {
		panic(err)
	}

	searchStreamRules, err := client.TweetSearchStreamRules(context.Background(), []twitter.TweetSearchStreamRuleID{})
	if err != nil {
		log.Panicf("tweet search stream rule callout error: %v", err)
	}

	enc, err := json.MarshalIndent(searchStreamRules, "", "    ")
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(string(enc))
}

func deleteRules(ids string) {
	client, err := getClient()
	if err != nil {
		panic(err)
	}

	ruleIDs := []twitter.TweetSearchStreamRuleID{}
	for _, id := range strings.Split(ids, ",") {
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

	fmt.Println(TWITTER_BEARER_TOKEN)
	client := &twitter.Client{
		Authorizer: authorize{
			Token: TWITTER_BEARER_TOKEN,
		},
		Client: http.DefaultClient,
		Host:   "https://api.twitter.com",
	}

	return client, nil
}
