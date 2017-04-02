package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	case today + 3:
		fmt.Println("In Three days.")
	case today + 4:
		fmt.Println("In Four days.")
	case today + 5:
		fmt.Println("In Five days.")
	case today + 6:
		fmt.Println("In Six days.")
	default:
		fmt.Println("Too far away.")
	}
}
