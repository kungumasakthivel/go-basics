package main

import "fmt"

func main() {
	fmt.Println("Pointers")

	var ptr *int
	fmt.Println(ptr)

	myNum := 23
	var pNum *int = &myNum
	fmt.Println(pNum)
	fmt.Println(*pNum)

	*pNum = *pNum + 7
	fmt.Println(*pNum)
}
