package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	tweetviz "github.com/ala24013/tweetviz/pkg"
)

func main() {
	sigs := make(chan os.Signal, 1)
	shutdown := make(chan int)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	tl := tweetviz.CreateTweetlist()
	go func() {
		fmt.Println("Starting webserver...")
		tweetviz.RunServer(tl)
	}()
	go func() {
		fmt.Println("Starting twitter stream...")
		tweetviz.Stream("Ian", tl, shutdown)
	}()
	<-sigs
	fmt.Println("Shutting down...")
}
