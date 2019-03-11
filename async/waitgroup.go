package main

import (
	"fmt"
	"sync"
	"time"
)

func waitGroup() {
	fmt.Println("[main] Starting")
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go service(&wg, i)
	}
	wg.Wait()
	fmt.Println("[main]Done")
}

func service(wg *sync.WaitGroup, id int) {
	time.Sleep(500 * time.Millisecond)
	fmt.Println("[thread] Servicing", id)
	wg.Done()
}
