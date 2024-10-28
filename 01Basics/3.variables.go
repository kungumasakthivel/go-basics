package main
import "fmt"

func main() {
	var num int16
	var str string
	var numDec float64
	var isval bool

	// default values if values are not specified
	fmt.Println(num) // 0
	fmt.Println(str) // ""
	fmt.Println(numDec)// 0
	fmt.Println(isval)// false

	// multiple assignment with specifying data type
	//declare multiple variables of the same type in the single declaration
	var v1, v2, v3 int = 10, 100, 1000
	fmt.Println(v1, v2, v3) 

	// multiple assignment without specifying data type
	// declare multiple variables of the different type in the single declaration
	var n1, s1, b1 = 22, "test", true
	fmt.Println(n1)
	fmt.Println(s1)
	fmt.Println(b1)

	// constant assignment
	// we cannot re-assign the same variable with different
	// values in future. Its value is fixed
	const PI = 3.14
	const name = "jon"
	fmt.Println(PI)
	// name = "peter" // you cannot re-assign the const variable (raise error)
	fmt.Println(name)
}