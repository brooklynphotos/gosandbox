package main

import (
	"fmt"
)

func sendEvents() {
	cmp := Component{
		listeners: make(map[string][]chan int),
	}
	clickCh1 := make(chan int)
	go func() {
		fmt.Println("Clicked 1:", <-clickCh1)
	}()
	clickCh2 := make(chan int)
	go func() {
		fmt.Println("Clicked 2:", <-clickCh2)
	}()
	scrollCh := make(chan int)
	go func() {
		fmt.Println("Scrolled", <-scrollCh)
	}()
	cmp.addListener("click", clickCh1)
	cmp.addListener("click", clickCh2)
	cmp.addListener("scroll", scrollCh)
	cmp.click()
	fmt.Scanln()
}

type callback func(num int)

type Component struct {
	listeners map[string][]chan int
}

func (cmp *Component) addListener(eventName string, ch chan int) {
	if cmp.listeners == nil {
		cmp.listeners = make(map[string][]chan int)
	}
	if _, ok := cmp.listeners[eventName]; ok {
		cmp.listeners[eventName] = append(cmp.listeners[eventName], ch)
	} else {
		cmp.listeners[eventName] = []chan int{ch}
	}
}

func (cmp *Component) click() {
	cmp.emit("click", 3)
}

func (cmp *Component) emit(eventName string, v int) {
	if chs, ok := cmp.listeners[eventName]; ok {
		for _, ch := range chs {
			ch <- v
		}
	}
}
