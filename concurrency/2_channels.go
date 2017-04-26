package main

import (
	"fmt"
	"time"
)

func sum(s []int, c chan int) {
	fmt.Println("Start summing", s)
	sum := 0
	for _, v := range s {
		fmt.Println(v)
		if v < 0 {
			time.Sleep(1000 * time.Millisecond)
		}
		sum += v
	}
	c <- sum // send sum to c
	fmt.Println("End summing", s)
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}
