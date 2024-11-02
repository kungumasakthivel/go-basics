package main

import "fmt"

func main() {
	fmt.Println("Loops")

	days := []string{"sunday", "monday", "tuesday", "wednesday"}
	fmt.Println(days)

	for _, day := range days {
		fmt.Println(day)
	}
	fmt.Println("---------")
	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}
	fmt.Println("---------")
	for i := range days {
		fmt.Println(days[i])
	}

	val := 1

	for val <= 10 {
		if val == 3 || val == 7 {
			val++
			continue
		}
		if val == 6 {
			goto lio
		}
		fmt.Println(val)
		if val == 8 {
			break
		}
		val++
	}
lio:
	fmt.Println("Hello")

}
