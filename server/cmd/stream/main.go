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
	go func() {
		fmt.Println("Starting twitter stream...")
		tweetviz.Stream("Trump", shutdown)
	}()
	<-sigs
	fmt.Println("Stopping stream...")
}
