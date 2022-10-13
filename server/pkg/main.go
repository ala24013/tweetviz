package tweetviz

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

var streamShutdown = make(chan int)

func Main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	tl := CreateTweetlist()
	go func() {
		fmt.Println("Starting webserver...")
		RunServer(tl)
	}()
	go func() {
		fmt.Println("Starting twitter stream...")
		Stream("Ian", tl)
	}()
	<-sigs
	fmt.Println("Shutting down...")
}
