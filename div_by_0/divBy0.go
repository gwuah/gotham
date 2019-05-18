package main

import "fmt"

func main() {
	c := 100 / 0.0
	// WONT COMIPILE
	fmt.Println(c)
}
