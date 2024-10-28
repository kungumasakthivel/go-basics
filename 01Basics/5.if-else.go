package main
import "fmt"

func main() {
	var v int = 100

	// if condition statement
	if v < 200 {
		fmt.Println("v is less than 200")
	}
	fmt.Println(v)


	var u int = 300
	
	// if-else condition statement
	if u < 200 {
		fmt.Println("u is less than 200")
	} else {
		fmt.Println("u is not less than 200")
	}
	fmt.Println(u)


	// nested if condition
	if u == 300 { // can also use brackets if(u == 300)
		if v == 100 { // can also use brackets if(v == 200)
			fmt.Println("u = 300, v = 100")
		}
	}

	// if else-if else condition statement
	if u > 300 {
		fmt.Println("u = 300")
	} else if v > 100 {
		fmt.Println("v = 100")
	} else {
		fmt.Println("not satisfied")
	}
}