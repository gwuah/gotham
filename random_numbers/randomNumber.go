package main

import (
	"fmt"
	"rand"
)

func main() {
	for i := 0; i < 10; i++ {
		var randomNumber int = rand.Int()
		fmt.Println(randomNumber)
	}
}
