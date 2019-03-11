package main

import (
	"fmt"
	"sync"
	"time"
)

func sendEachOther() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			ch <- i
		}
		close(ch)
	}()
	for x := range ch {
		fmt.Println(x)
	}
	// reversed, main goroutine sends data
	ch2 := make(chan int)
	go func(c chan int) {
		fmt.Println("Inside got", <-c, "!")
	}(ch2)
	ch2 <- 3

	fmt.Println("Go")
}

func square(c chan int) {
	fmt.Println("[square] starting")
	num := <-c
	c <- num * num
	fmt.Println("[square] done")
}

func cube(c chan int) {
	fmt.Println("[cube] starting")
	num := <-c
	c <- num * num * num
	fmt.Println("[cube] done")
}
func twoChannels() {
	sc := make(chan int)
	cc := make(chan int)

	go square(sc)
	go cube(cc)

	fmt.Println("Starting with a number")
	fmt.Println("Feeding square channel")
	num := 3
	sc <- num
	fmt.Println("Resumed")

	fmt.Println("Feeding cube channel")
	cc <- num
	fmt.Println("resumed")

	fmt.Println("Reading from channels")
	fmt.Println("Got square:", <-sc)
	fmt.Println("Got cube:", <-cc)
}

func fetchData(c chan int, x int) {
	fmt.Println("Fetching for id=", x)
	time.Sleep(time.Duration(x*100) * time.Millisecond)
	c <- x
}

func selectChannels() {
	c1 := make(chan int)
	c2 := make(chan int)

	go fetchData(c1, 2)
	go fetchData(c2, 110)

	select {
	case res := <-c1:
		fmt.Println("[main] Got c1", res)
	case res := <-c2:
		fmt.Println("[main] Got c2", res)
	}
	fmt.Println("[main] Done!")
}

func waitOnWorkers() {
	var wg sync.WaitGroup
	fmt.Println("Starting workers")
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go fetchRemoteData(&wg, i)
	}
	wg.Wait()
	fmt.Println("Done for all workers")
}

func fetchRemoteData(wg *sync.WaitGroup, instance int) {
	fmt.Println("Starting the fetch for", instance)
	time.Sleep(1 * time.Second)
	fmt.Println("Done waiting for", instance)
	wg.Done()
}
