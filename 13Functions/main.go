package main

import "fmt"

// main act as entry point for program
// not allowed to create function inside function
// function signature - func add(n1 int, n2 int) int {} - int outside bracket is return type should be specified
func main() {
	fmt.Println("Function")
	greet()
	result := add(34, 87)
	fmt.Println(result)
	result2 := proAdd(2, 4, 5, 8, 2)
	fmt.Println(result2)
	fmt.Println(returnTwoType("Manish", 20202))
}

func add(n1 int, n2 int) int {
	return n1 + n2
}

func proAdd(values ...int) int { // acts like overloading
	total := 0

	for _, val := range values {
		total += val
	}
	return total
}

func greet() {
	fmt.Println("Hello")
}

func returnTwoType(name string, rollno int) (string, int) {
	return "Name: " + name, rollno
}

// func ananomous() {
// 	fmt.Println("ananomous function")
// }
