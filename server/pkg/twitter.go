package tweetviz

import (
	"flag"
	"fmt"
	"log"

	"github.com/coreos/pkg/flagutil"
	"github.com/dghubble/go-twitter/twitter"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

var (
	TWITTER_KEY    string
	TWITTER_SECRET string
)

func Stream(query string, shutdown <-chan int) {
	client, err := getClient()
	if err != nil {
		panic(err)
	}

	demux := twitter.NewSwitchDemux()
	demux.Tweet = func(tweet *twitter.Tweet) {
		fmt.Println(tweet.Text)
	}
	demux.DM = func(dm *twitter.DirectMessage) {
		fmt.Println(dm.SenderID)
	}
	demux.Event = func(event *twitter.Event) {
		fmt.Printf("%#v\n", event)
	}

	params := &twitter.StreamFilterParams{
		Track:         []string{query},
		StallWarnings: twitter.Bool(true),
	}
	stream, err := client.Streams.Filter(params)
	if err != nil {
		panic(err)
	}

	go demux.HandleChan(stream.Messages)

	<-shutdown
	stream.Stop()
}

func getClient() (*twitter.Client, error) {
	if TWITTER_KEY == "" {
		flag.StringVar(&TWITTER_KEY, "KEY", "", "Twitter Key")
	}
	if TWITTER_SECRET == "" {
		flag.StringVar(&TWITTER_SECRET, "SECRET", "", "Twitter Secret")
	}
	flag.Parse()
	flagutil.SetFlagsFromEnv(flag.CommandLine, "TWITTER")

	if TWITTER_KEY == "" || TWITTER_SECRET == "" {
		log.Fatal("Application Access Token required")
	}

	// oauth2 configures a client that uses app credentials to keep a fresh token
	config := &clientcredentials.Config{
		ClientID:     TWITTER_KEY,
		ClientSecret: TWITTER_SECRET,
		TokenURL:     "https://api.twitter.com/oauth2/token",
	}
	// http.Client will automatically authorize Requests
	httpClient := config.Client(oauth2.NoContext)

	client := twitter.NewClient(httpClient)

	return client, nil
}
