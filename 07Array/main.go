package main

import "fmt"

func main() {
	fmt.Println("Arrays")
	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Banana"
	fruitList[3] = "Carrot"
	fmt.Println(fruitList)
	fmt.Println(len(fruitList))

	var veg = [3]string{"patato", "beans", "mushroom"}
	fmt.Println(veg)
}
