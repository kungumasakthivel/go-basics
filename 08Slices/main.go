package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Slices")

	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("%T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:4])
	fmt.Println(fruitList)

	highScore := make([]int, 4)
	highScore[0] = 234
	highScore[1] = 678
	highScore[2] = 456
	highScore[3] = 567
	//highScore[4] = 678 // you can't do this but can use append
	highScore = append(highScore, 789, 890)
	fmt.Println(highScore)

	sort.Ints(highScore)
	fmt.Println(highScore)
	fmt.Println(sort.IntsAreSorted(highScore))
	fmt.Println("--------------------------")
	// remove value based on idx

	var alpha = []string{"a", "b", "c", "d", "e"}
	fmt.Println(alpha)
	var idx int = 2
	alpha = append(alpha[:idx], alpha[idx+1:]...)
	fmt.Println(alpha)
}
