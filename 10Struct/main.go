package main

import "fmt"

func main() {
	fmt.Println("struct golang")
	// no inheritance in golang; no super or parent class

	x := User{"x", "x@gmail.com", true, 100}
	fmt.Println(x)
	fmt.Printf("Details are: %+v\n", x)
	fmt.Println(x.Name)
	fmt.Println(x.Age)
	fmt.Printf("Name is %v and email is %v\n", x.Name, x.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
