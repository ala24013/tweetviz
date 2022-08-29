package main

import (
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go server.runServer(wg)
	go server.stream(wg)
	wg.Wait()
}
