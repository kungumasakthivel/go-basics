package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://jsonplaceholder.typicode.com/posts/1"

func main() {
	fmt.Println("WEB request")

	res, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	fmt.Println(res)
	fmt.Printf("%T", res)
	defer res.Body.Close()

	databyte, err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(databyte))
}
