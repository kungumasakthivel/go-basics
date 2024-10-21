package main

import "fmt"

func main() {
	var x int16 = 170
	var y int16 = 83

	fmt.Println("Add", x+y)
	fmt.Println("Sub", x-y)
	fmt.Println("Mul", x*y)
	fmt.Println("Div", x/y)
	fmt.Println("Mod", x%y)

	var a float32 = 20.45
	var b float32 = 10.55

	fmt.Println(a-b);

	var ac complex128 = complex(6, 2)
	var bc complex64 = complex(9, 2)

	fmt.Println(ac)
	fmt.Println(bc)

	var real = real(ac)
	var imag = imag(ac)
	fmt.Println("Real num", real, "Imaginary num",imag)


	var str1 = "GfG"
	var str2 = "gFg"
	var str3 = "GfG"
	var r1 = str1 == str2
	var r2 = str1 == str3

	fmt.Println(r1, r2)
}