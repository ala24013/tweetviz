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
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		fmt.Println("Starting webserver...")
		tweetviz.RunServer()
	}()
	go func() {
		fmt.Println("Starting twitter stream...")
		tweetviz.Stream("test")
	}()
	<-sigs
	fmt.Println("Shutting down...")
}
