package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Files")
	content := "This needs to go in a file"

	file, err := os.Create("./myfile.txt")

	if err != nil {
		panic(err) // panic shutdown execution of program and show error
	}

	length, err := io.WriteString(file, content)

	if err != nil {
		panic(err)
	}

	fmt.Println(length)

	file.Close()

	readFile("./myfile.txt")
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}
	fmt.Println(databyte)
	fmt.Println(string(databyte))
}
