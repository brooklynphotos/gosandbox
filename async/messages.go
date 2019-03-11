package main

import "fmt"

func sendMessage() {
	ch := make(chan int)
	makeAccount(ch)
	ch <- 3
	fmt.Scanln()
}

func makeAccount(ch chan int) {
	acct := Account{}
	go func(ch chan int) {
		acct.processPayment(<-ch)
	}(ch)
}

type Account struct{}

func (account *Account) processPayment(amount int) {
	fmt.Println("Processing ", amount)
}
