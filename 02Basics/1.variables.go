package main

import "fmt"

const LoginId string = "654567drtyujhgf"

func main() {
	var username string = "manish"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)
	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)
	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)
	var smallFloat float32 = 255.123456789
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)
	anotherWayToAssignVar := 20
	fmt.Println(anotherWayToAssignVar)
	fmt.Println(LoginId)
	fmt.Printf("Variable is of type: %T \n", LoginId)
}
