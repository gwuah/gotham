package main

import (
	"fmt"
	"math"
)

// func Uint8FromInt(n int) (uint8, error) {

// }

func main() {
	// var a float64
	var b int32

	whole, fraction := math.Modf(4.7)

	// a = 12
	fmt.Println(whole, fraction)
	// b = a + a
	b = b + 5
	fmt.Println(b)
}
