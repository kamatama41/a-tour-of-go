package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	n0, n1 := 1, 0
	return func() int {
		newN0, newN1 := n1, n0 + n1
		n0, n1 = newN0, newN1
		return n1
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
