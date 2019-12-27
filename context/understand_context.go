/*
 * Based on this http://p.agnihotry.com/post/understanding_the_context_package_in_golang/
 */

package main

import "fmt"

func printHello(ch chan int) {
	fmt.Println("Hello from printHello")
	ch <- 2
}

func run() {
	ch := make(chan int)
	go func() {
		fmt.Println("Local hello")
		ch <- 1
	}()
	go printHello(ch)
	fmt.Println("From main")
	sig := <-ch
	fmt.Println("Got signal ", sig)
	fmt.Println("Got another signal ", <-ch)
}
