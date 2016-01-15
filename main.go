package main

import "fmt"

func main() {
	fmt.Println("\033[2J")
	// fmt.Println("\033[?5h")
	fmt.Println("Name:", levels[1].name)
	fmt.Println("Bottom", levels[1].layout[18])
}
