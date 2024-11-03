package main

import "fmt"

func main() {
	defer fmt.Println("Hello 2") // execute at last befor last main curly braces
	defer fmt.Println("3")
	defer fmt.Println("4") // 2 or more defer - executation last in first out - stack principle
	fmt.Println("Hello, world!")
	myDefer()
}

func myDefer() {
	defer fmt.Println()
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}
