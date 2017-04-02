package main

import (
	"fmt"
	"math"
)

//func Sqrt(x float64) float64 {
//	z := float64(1)
//	for i := 0; i < 10; i++ {
//		z = z - (z * z - x) / (2 * z)
//		fmt.Println(z)
//	}
//	return z
//}


func Sqrt(x float64) float64 {
	z1 := float64(1)
	z2 := float64(100)
	cnt := 0
	for z1 != z2 {
		var tmp = z2
		z2 = z1 - (z1 * z1 - x) / (2 * z1)
		z1 = tmp
		cnt++
		fmt.Println(z1, z2, cnt)
	}
	return z2
}

func main() {
	fmt.Println(Sqrt(3))
	fmt.Println(math.Sqrt(3))
}
