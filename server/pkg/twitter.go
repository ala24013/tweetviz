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

func Stream(query string) {
	client, err := getClient()
	if err != nil {
		panic(err)
	}

	params := &twitter.StreamFilterParams{
		Track:         []string{"kitten"},
		StallWarnings: twitter.Bool(true),
	}
	stream, err := client.Streams.Filter(params)
	if err != nil {
		panic(err)
	}

	for m := range stream.Messages {
		fmt.Println(m)
	}
}

func getClient() (*twitter.Client, error) {
	flags := struct {
		twitterKey    string
		twitterSecret string
	}{}

	flag.StringVar(&flags.twitterKey, "KEY", "", "Twitter Key")
	flag.StringVar(&flags.twitterSecret, "SECRET", "", "Twitter Secret")
	flag.Parse()
	flagutil.SetFlagsFromEnv(flag.CommandLine, "TWITTER")

	if flags.twitterKey == "" || flags.twitterSecret == "" {
		log.Fatal("Application Access Token required")
	}

	// oauth2 configures a client that uses app credentials to keep a fresh token
	config := &clientcredentials.Config{
		ClientID:     flags.twitterKey,
		ClientSecret: flags.twitterSecret,
		TokenURL:     "https://api.twitter.com/oauth2/token",
	}
	// http.Client will automatically authorize Requests
	httpClient := config.Client(oauth2.NoContext)

	client := twitter.NewClient(httpClient)

	return client, nil
}
