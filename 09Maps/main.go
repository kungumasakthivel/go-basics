package main

import "fmt"

func main() {
	fmt.Println("Maps")

	language := make(map[int]string)

	language[1] = "JS"
	language[2] = "CSS"
	language[3] = "HTML"

	fmt.Println(language)
	fmt.Println(language[1])
	fmt.Println(language[3])

	delete(language, 2)
	fmt.Println(language)

	for key, value := range language {
		fmt.Printf("key %v : value %v\n", key, value)
	}
}
