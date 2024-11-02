package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch Case")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of dice:", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1, you can start")
	case 2:
		fmt.Println(2)
	case 3:
		fmt.Println(3)
	case 4:
		fmt.Println(4)
		// fallthrough if 4 in dice then 5 also printed
	case 5:
		fmt.Println(5)
	case 6:
		fmt.Println("Dice value is 6 and roll dice again")
	default:
		fmt.Println("Above my knowledge")
	}
}
