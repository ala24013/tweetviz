package main

import (
	"sync"

	"github.com/ala24013/tweetviz"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go tweetviz.RunServer(wg)
	go tweetviz.Stream(wg)
	wg.Wait()
}
