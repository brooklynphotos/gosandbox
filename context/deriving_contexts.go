/**
 * Again, based on http://p.agnihotry.com/post/understanding_the_context_package_in_golang/
 * code is https://github.com/pagnihotry/golang_samples/blob/master/go_context_sample.go
 */

package main

import (
	"context"
	"fmt"
	"time"
)

// MainSleep is Maximum time main will sleep for
const MainSleep = 2

// DoWorkTimeout is duration of Do work timeout context
const DoWorkTimeout = 4

// SleepContextDuration is the sleep time the sleepContext function uses for the sleep function
const SleepContextDuration = 15

func sleepFor(seconds int, ch chan int) {
	defer func() {
		fmt.Println("Sleep for done")
	}()

	fmt.Println("Start sleeping for ", seconds, " seconds")
	time.Sleep(time.Duration(seconds) * time.Second)
	fmt.Println("Done sleeping")

	if ch != nil {
		ch <- seconds
	}
}

func sleepContext(ctx context.Context, ch chan bool) {
	// always cleanup and send ch a true signal
	defer func() {
		fmt.Println("Done sleeping with context")
		ch <- true
	}()

	// channel for the sleep function
	sleepCh := make(chan int)
	go sleepFor(SleepContextDuration, sleepCh)

	select {
	case <-ctx.Done():
		// this ctx came from another derived ctx, so Done() returns when any of the contexts along the hierarchy is canceled
		fmt.Println("Sleepcontext is notified to return")
	case sleepTime := <-sleepCh:
		// we made it in time, before any context was canceled
		fmt.Println("Sleep for ", sleepTime, " seconds")
	}
}

func doWork(ctx context.Context) {
	timeoutCtx, cancelFunction := context.WithTimeout(ctx, time.Duration(DoWorkTimeout)*time.Second)

	// when all is done, make sure we cancel the context so all users of timeoutCtx knows if they want to
	defer func() {
		fmt.Println("Do work completed")
		cancelFunction()
	}()

	ch := make(chan bool)
	go sleepContext(timeoutCtx, ch)

	select {
	case <-ctx.Done():
		// We are notified that the passed in context is done
		fmt.Println("Time's up for doWork, leave")
	case <-ch:
		fmt.Println("Channel came back after being put to sleep!")
	}
}

func main() {
	mainCtx := context.Background()
	// cancel context derived from background
	cancelCtx, cancelFunc := context.WithCancel(mainCtx)

	// clean up preparation
	defer func() {
		fmt.Println("Cleaning up")
		cancelFunc()
	}()

	// Will sleep for some time, and if the remaining main doesn't finish, cancel kicks in
	go func() {
		sleepFor(MainSleep, nil)
		cancelFunc()
		fmt.Println("Done sleeping, canceling context")
	}()

	// what we do while the main thread sleeps
	doWork(cancelCtx)
}
