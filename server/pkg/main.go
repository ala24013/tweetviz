package tweetviz

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

var streamShutdown = make(chan int)

// Main runs tweetviz, instantiating the initial stream and the webserver
func Main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	tl := CreateTweetlist()
	go func() {
		fmt.Println("Starting webserver...")
		RunServer(tl)
	}()
	<-sigs
	Shutdown()
}

// Shutdown shuts down the various running go routines
func Shutdown() {
	fmt.Println("Shutting down...")
	<-shutdownWs
	<-streamShutdown
}
