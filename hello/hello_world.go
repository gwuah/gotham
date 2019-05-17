package main

import (
	"fmt"
	sys "fmt"
	"os"
)

func firstImport() {
	fmt.Println("Hello World from Griffith")
}

func getEnv() {
	goos := os.Getenv("GOOS")
	path := os.Getenv("PATH")
	fmt.Printf("The path name is: %s\n", path)
	fmt.Printf("The os name is: %s\n", goos)
}

func main() {
	firstImport()
	getEnv()
	sys.Println("Second Hello")
}
