package main

import "fmt"

var a = "String"

func n() {
	fmt.Println(a)
}

func m() {
	a := 40
	fmt.Println(a)
}

func main() {
	n()
	m()
	n()
}
