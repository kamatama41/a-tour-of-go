package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 1)
	go func() {
		time.Sleep(500 * time.Millisecond)
		ch <- 1
	}()
	go func() {
		time.Sleep(1000 * time.Millisecond)
		ch <- 2
	}()
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
