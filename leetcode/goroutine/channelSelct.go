package main

import (
	"fmt"
)

func main() {
	ch1, ch2 := make(chan int), make(chan int)
	go func(ch chan int) {
		for i := 0; ; i++ {
			if i%2 == 0 {
				ch <- i
			}
		}
	}(ch1)
	go func(ch chan int) {
		for i := 0; ; i++ {
			if i%2 != 0 {
				ch <- i
			}
		}
	}(ch2)
	for {
		select {
		case u := <-ch1:
			fmt.Println("ch1: %d", u)
		case v := <-ch2:
			fmt.Println("ch2: %d", v)
		}
	}
}
