package main

import (
	"fmt"
	"sync"
	"time"
)

func mapReduce() {
	workerCount := 3
	taskCount := 5
	bufferSize := 10
	mapChan := make(chan int, bufferSize)
	reduceChan := make(chan int, bufferSize)
	var wg sync.WaitGroup

	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go mapper(&wg, mapChan, reduceChan, i)
	}

	for i := 0; i < taskCount; i++ {
		mapChan <- i
	}
	fmt.Printf("[main] Wrote %d tasks\n", taskCount)
	time.Sleep(time.Millisecond)
	close(mapChan)

	wg.Wait()

	sum := 0
	for i := 0; i < taskCount; i++ {
		sum += <-reduceChan
	}
	fmt.Println("[main]Done with sum", sum)
}

func mapper(wg *sync.WaitGroup, mapChan chan int, reduceChan chan int, id int) {
	for x := range mapChan {
		time.Sleep(300 * time.Millisecond)
		fmt.Printf("[worker %d] Doing the mapping work\n", id)
		reduceChan <- x * x
	}
	wg.Done()
}
